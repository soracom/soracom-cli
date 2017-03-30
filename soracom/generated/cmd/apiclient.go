package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
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
	hc := http.DefaultClient

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
	method      string
	path        string
	query       string
	contentType string
	body        string
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

	res, err := ac.httpClient.Do(req)
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
