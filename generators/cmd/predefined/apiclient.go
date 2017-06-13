package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
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

	var endpoint = getSpecifiedEndpoint()

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
}

func (ac *apiClient) callAPI(params *apiParams) (string, error) {
	url := ac.endpoint + ac.basePath + params.path
	if params.query != "" {
		url += "?" + params.query
	}
	//fmt.Printf("url == %v\n", url)
	req, err := http.NewRequest(params.method, url, strings.NewReader(params.body))
	if err != nil {
		return "", err
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
		return "", err
	}
	defer res.Body.Close()

	if ac.verbose {
		dumpHTTPResponse(res)
		fmt.Println("==========")
	}

	if res.StatusCode >= http.StatusBadRequest {
		return "", newAPIError(res)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return bytes.NewBuffer(b).String(), nil
}

func (ac *apiClient) doHTTPRequestWithRetries(req *http.Request) (*http.Response, error) {
	var err error
	backoffSeconds := []int{10, 10, 20, 30, 50}
	for _, wait := range backoffSeconds {
		var res *http.Response
		res, err = ac.httpClient.Do(req)
		if err == nil && !retryableError(res.StatusCode) {
			return res, nil
		}
		if err != nil && res != nil && res.Body != nil {
			defer res.Body.Close()
		}

		ac.reportWaitingBeforeRetrying(res, err, wait)
		time.Sleep(time.Duration(wait) * time.Second)
		ac.reportRetrying()
	}

	return nil, err
}

func (ac *apiClient) reportWaitingBeforeRetrying(res *http.Response, err error, wait int) {
	if !ac.verbose {
		return
	}
	fmt.Fprintf(os.Stderr, "error detected. ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "http status code == %d\n", res.StatusCode)
	}
	fmt.Fprintf(os.Stderr, "wait for %d seconds ...\n", wait)
}

func (ac *apiClient) reportRetrying() {
	if !ac.verbose {
		return
	}
	fmt.Fprintf(os.Stderr, "trying it again\n")
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
		fmt.Println(err)
		return
	}
	fmt.Println(string(dump))
}

func dumpHTTPResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(dump))
}
