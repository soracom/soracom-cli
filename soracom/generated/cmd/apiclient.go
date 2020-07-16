package cmd

import (
	"bytes"
	"encoding/json"
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

	"github.com/pkg/errors"
	"github.com/soracom/soracom-cli/generators/lib"
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

func newAPIError(respBody string) *apiError {
	return &apiError{
		ResponseBody: respBody,
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
	method                            string
	path                              string
	query                             url.Values
	contentType                       string
	body                              string
	noRetryOnError                    bool
	noVersionCheck                    bool
	doPagination                      bool
	paginationKeyHeaderInResponse     string
	paginationRequestParameterInQuery string
}

// params.path and params.query must be escaped before calling this func
func (ac *apiClient) callAPI(params *apiParams) (string, error) {
	var (
		resBody       string
		latestVersion string
	)

	for {
		u, err := ac.constructURL(params)
		if err != nil {
			return "", err
		}

		req, err := ac.constructRequest(u, params)
		if err != nil {
			return "", err
		}

		if ac.verbose {
			dumpHTTPRequest(req)
		}

		res, rb, err := ac.doRequest(req, params)
		if err != nil {
			return "", err
		}
		if ac.verbose && res != nil {
			dumpHTTPResponse(res)
		}

		if res.StatusCode >= http.StatusBadRequest {
			return "", newAPIError(rb)
		}
		latestVersion = res.Header.Get("x-soracom-cli-version")

		if !params.doPagination {
			resBody = rb
			break
		}

		resBody, err = concatJSONArray(resBody, rb)
		if err != nil {
			return "", err
		}

		k, v := getPaginationKeyValue(res, params)
		if k == "" || v == "" {
			break
		}

		setPaginationKeyValue(params, v)
	}

	if ac.verbose {
		lib.PrintfStderr("==========\n")
	}

	if !params.noVersionCheck {
		if isNewerThanCurrentVersion(latestVersion) {
			lib.PrintfStderr(TRCLI("cli.new-version-is-released"), latestVersion, version)
		}
	}

	return resBody, nil
}

// arr1 = "[1,2]"
// arr2 = "[3]"
// returns "[1,2,3]"
func concatJSONArray(arr1, arr2 string) (string, error) {
	if arr1 == "" {
		return arr2, nil
	}

	a1 := make([]interface{}, 0)
	a2 := make([]interface{}, 0)

	err := json.Unmarshal([]byte(arr1), &a1)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(arr2), &a2)
	if err != nil {
		return "", err
	}

	a := append(a1, a2...)

	b, err := json.Marshal(a)
	return string(b), err
}

func (ac *apiClient) constructURL(params *apiParams) (*url.URL, error) {
	urlString := ac.endpoint + ac.basePath + params.path
	if params.query != nil {
		urlString += "?" + params.query.Encode()
	}
	return url.Parse(urlString)
}

func (ac *apiClient) constructRequest(u *url.URL, params *apiParams) (*http.Request, error) {
	req, err := http.NewRequest(params.method, u.String(), strings.NewReader(params.body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", fmt.Sprintf("soracom-cli/%s", version))

	if params.contentType != "" {
		req.Header.Set("Content-Type", params.contentType)
		if params.contentType == "application/octet-stream" {
			req.Header.Set("Content-Length", fmt.Sprintf("%d", len(params.body)))
		}
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

	return req, nil
}

func (ac *apiClient) doRequest(req *http.Request, params *apiParams) (*http.Response, string, error) {
	var (
		res *http.Response
		err error
	)

	if params.noRetryOnError {
		res, err = ac.httpClient.Do(req)
	} else {
		res, err = ac.doHTTPRequestWithRetries(req, params)
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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}
	return res, bytes.NewBuffer(b).String(), nil
}

func getPaginationKeyValue(res *http.Response, params *apiParams) (string, string) {
	k := params.paginationKeyHeaderInResponse
	if k != "" {
		v := res.Header.Get(k)
		return k, v
	}
	return "", ""
}

func setPaginationKeyValue(params *apiParams, v string) {
	k := params.paginationRequestParameterInQuery
	if k != "" {
		params.query.Set(k, v)
	}
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

func (ac *apiClient) doHTTPRequestWithRetries(req *http.Request, params *apiParams) (*http.Response, error) {
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
		req.Body = ioutil.NopCloser(strings.NewReader(params.body)) // reload body contents
	}

	return nil, errors.New("unable to receive successful response with some retires")
}

func (ac *apiClient) reportWaitingBeforeRetrying(res *http.Response, err error, wait int) {
	if !ac.verbose {
		return
	}
	lib.PrintfStderr("error detected. ")
	if err != nil {
		lib.PrintfStderr("%+v\n", err)
	} else {
		lib.PrintfStderr("http status code == %d\n", res.StatusCode)
		dumpHTTPResponse(res)
	}
	lib.PrintfStderr("wait for %d seconds ...\n", wait)
}

func (ac *apiClient) reportRetrying() {
	if !ac.verbose {
		return
	}
	lib.PrintfStderr("trying it again\n")
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
	dumpBody := req.Header.Get("Content-Type") != "application/octet-stream"
	dump, err := httputil.DumpRequest(req, dumpBody)
	if err != nil {
		lib.PrintfStderr("error while dumping http request header and body: %s\n", err)
		return
	}
	lib.PrintfStderr("%s\n", string(dump))
}

func dumpHTTPResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		lib.PrintfStderr("error while dumping http response header: %s\n", err)
		return
	}
	lib.PrintfStderr("%s\n", string(dump))
}
