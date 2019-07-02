package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// APIClient provides an access to SORACOM REST API
type apiClient struct {
	httpClient *http.Client
	APIKey     string
	Token      string
	OperatorID string
	endpoint   string
	basePath   string
	language   string
	verbose    bool
}

// APIError represents an error occurred while calling API
type apiError struct {
	ResponseBody string
}

func newAPIError(resp *http.Response) *apiError {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &apiError{
			ResponseBody: err.Error(),
		}
	}
	return &apiError{
		ResponseBody: string(body),
	}
}

func (ae *apiError) Error() string {
	return ae.ResponseBody
}

type apiClientOptions struct {
	BasePath string
	Language string
	Endpoint string
}

// New creates an instance of APIClient
func newAPIClient(options *apiClientOptions) *apiClient {
	hc := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			ResponseHeaderTimeout: 120 * time.Second,
		},
		Timeout: 120 * time.Second,
	}

	var endpoint string
	if options.Endpoint != "" {
		endpoint = options.Endpoint
	} else {
		endpoint = getSpecifiedEndpoint()
	}

	var basePath = "/"
	if options != nil && options.BasePath != "" {
		basePath = options.BasePath
	}

	var language = "en"
	if options != nil && options.Language != "" {
		language = options.Language
	}

	var verbose = false
	v := os.Getenv("SORACOM_VERBOSE")
	if v == "1" {
		verbose = true
	}

	return &apiClient{
		httpClient: hc,
		APIKey:     "",
		Token:      "",
		OperatorID: "",
		endpoint:   endpoint,
		basePath:   basePath,
		language:   language,
		verbose:    verbose,
	}
}

type apiParams struct {
	method         string
	path           string
	query          string
	contentType    string
	body           string
	noRetryOnError bool
	noVersionCheck bool
}

// params.path and params.query must be escaped before calling this func
func (ac *apiClient) callAPI(params *apiParams) (http.Header, string, error) {
	urlString := ac.endpoint + ac.basePath + params.path
	if params.query != "" {
		urlString += "?" + params.query
	}
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, "", err
	}

	req, err := http.NewRequest(params.method, u.String(), strings.NewReader(params.body))
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("User-Agent", fmt.Sprintf("soracom-cli/%s", version))

	if params.contentType != "" {
		req.Header.Set("Content-Type", params.contentType)
	}

	if ac.APIKey != "" {
		req.Header.Set("X-Soracom-API-Key", ac.APIKey)
	}

	if ac.Token != "" {
		req.Header.Set("X-Soracom-Token", ac.Token)
	}

	if ac.language != "" {
		req.Header.Set("X-Soracom-Lang", ac.language)
	}

	if ac.verbose {
		dumpHTTPRequest(req)
	}

	var res *http.Response
	if params.noRetryOnError {
		res, err = ac.httpClient.Do(req)
	} else {
		res, err = ac.doHTTPRequestWithRetries(req)
	}
	if err != nil {
		return nil, "", err
	}
	if res == nil {
		return nil, "", errors.New("nil response received")
	}
	defer func() {
		// #nosec G104
		res.Body.Close()
		// #nosec G104
		io.Copy(ioutil.Discard, res.Body)
	}()

	if ac.verbose {
		dumpHTTPResponse(res)
		printfStderr("==========\n")
	}

	if res.StatusCode >= http.StatusBadRequest {
		return nil, "", newAPIError(res)
	}

	if !params.noVersionCheck {
		latestVersion := res.Header.Get("x-soracom-cli-version")
		if isNewerThanCurrentVersion(latestVersion) {
			printfStderr(TRCLI("cli.new-version-is-released"), latestVersion, version)
		}
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}
	return res.Header, bytes.NewBuffer(b).String(), nil
}

// version strings are in the form of "v1.22.333" or "v0.0.1"
func isNewerThanCurrentVersion(latestVersion string) bool {
	cv := versionInt(version)
	lv := versionInt(latestVersion)
	return cv < lv
}

func versionInt(ver string) uint32 {
	s := splitVersionString(ver)
	if len(s) < 3 {
		return 0
	}

	var n uint32
	shift := uint(24)
	for i := 0; i < 4; i++ {
		if len(s) <= i {
			break
		}
		x, err := strconv.Atoi(s[i])
		if err == nil {
			n |= uint32((x & 0xff) << shift)
		}
		shift -= 8
	}
	return n
}

var versionStringRegexp = regexp.MustCompile("([[:digit:]]+)[[:^digit:]]*")

func splitVersionString(ver string) []string {
	m := versionStringRegexp.FindAllStringSubmatch(ver, -1)
	if len(m) < 2 {
		return []string{}
	}
	result := make([]string, len(m))
	for i, s := range m {
		result[i] = s[1]
	}
	return result
}

func (ac *apiClient) doHTTPRequestWithRetries(req *http.Request) (*http.Response, error) {
	backoffSeconds := []int{10, 10, 20, 30, 50}
	for _, wait := range backoffSeconds {
		res, err := ac.httpClient.Do(req)
		if err == nil && !retryableError(res.StatusCode) {
			return res, nil
		}
		if err != nil && res != nil && res.Body != nil {
			defer func() {
				io.Copy(ioutil.Discard, res.Body)
				res.Body.Close()
			}()
		}

		ac.reportWaitingBeforeRetrying(res, err, wait)
		time.Sleep(time.Duration(wait) * time.Second)
		ac.reportRetrying()
	}

	return nil, errors.New("unable to receive successful response with some retires")
}

func (ac *apiClient) reportWaitingBeforeRetrying(res *http.Response, err error, wait int) {
	if !ac.verbose {
		return
	}
	printfStderr("error detected. ")
	if err != nil {
		printfStderr("%+v\n", err)
	} else {
		printfStderr("http status code == %d\n", res.StatusCode)
	}
	printfStderr("wait for %d seconds ...\n", wait)
}

func (ac *apiClient) reportRetrying() {
	if !ac.verbose {
		return
	}
	printfStderr("trying it again\n")
}

func retryableError(httpStatus int) bool {
	if httpStatus == http.StatusTooManyRequests {
		return true
	}
	if httpStatus < http.StatusInternalServerError {
		return false
	}

	return true
}

// SetVerbose sets if verbose output is enabled or not
func (ac *apiClient) SetVerbose(verbose bool) {
	ac.verbose = verbose
}

func dumpHTTPRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		printfStderr("%s\n", err)
		return
	}
	printfStderr("%s\n", string(dump))
}

func dumpHTTPResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		printfStderr("%s\n", err)
		return
	}
	printfStderr("%s\n", string(dump))
}

func printfStderr(format string, args ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, format, args...)
	if err != nil {
		//fmt.Printf("err: %+v\n", err) // this messes up stdout
	}
}
