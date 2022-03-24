package cmd

import (
	"net/http"
	"os"
	"testing"

	"github.com/elazarl/goproxy"
	"github.com/tj/assert"
)

type footprints struct {
	http  int
	https int
}

func setupProxyServerForTest(proxyAddr string, fp *footprints) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.Tr = &http.Transport{} // To avoid using this proxy to access this proxy itself

	proxy.OnRequest(goproxy.DstHostIs("example.com")).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		fp.http++
		return req, nil
	})

	proxy.OnRequest().HandleConnect(goproxy.FuncHttpsHandler(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		fp.https++
		return goproxy.RejectConnect, host
	}))

	go http.ListenAndServe(proxyAddr, proxy)
}

func TestCallAPIWithProxy(t *testing.T) {
	origEnvVars := saveEnvVars([]string{"HTTP_PROXY"})
	defer restoreEnvVars(origEnvVars)

	proxyPort := ":18080"
	proxyAddr := "http://localhost" + proxyPort
	err := os.Setenv("HTTP_PROXY", proxyAddr)
	assert.NoError(t, err)
	err = os.Setenv("HTTPS_PROXY", proxyAddr)
	assert.NoError(t, err)

	fp := &footprints{}
	setupProxyServerForTest(proxyPort, fp)

	// Test HTTP
	ac := newAPIClient(&apiClientOptions{Endpoint: "http://example.com"})
	// Ignoring response and error because example.com always returns 404 Not Found for the path 'v1/subscribers'
	_, _ = ac.callAPI(&apiParams{
		method:         "GET",
		path:           "v1/subscribers",
		noRetryOnError: true,
	})
	assert.Equal(t, 1, fp.http, "proxy server should be accessed when HTTP_PROXY env var is set")

	// Test HTTPS
	ac = newAPIClient(&apiClientOptions{Endpoint: "https://api.soracom.io"})
	// Ignoring response and error because api.soracom.io always returns 400 Bad request with the error message 'invalid api key'
	_, _ = ac.callAPI(&apiParams{
		method:         "GET",
		path:           "v1/subscribers",
		noRetryOnError: true,
	})
	assert.Equal(t, 1, fp.https, "proxy server should be accessed when HTTP_PROXY env var is set")

}

func saveEnvVars(vars []string) map[string]string {
	result := make(map[string]string)
	for _, v := range vars {
		result[v] = os.Getenv(v)
	}
	return result
}

func restoreEnvVars(origVars map[string]string) {
	for k, v := range origVars {
		os.Setenv(k, v)
	}
}

func TestConcatJSONArray(t *testing.T) {
	var testData = []struct {
		Name     string
		Arr1     string
		Arr2     string
		Expected string
	}{
		{
			Name:     "pattern 1",
			Arr1:     "[1,2]",
			Arr2:     "[3]",
			Expected: "[1,2,3]",
		},
		{
			Name:     "pattern 2",
			Arr1:     "[]",
			Arr2:     "[1]",
			Expected: "[1]",
		},
		{
			Name:     "pattern 3",
			Arr1:     "[1]",
			Arr2:     "[]",
			Expected: "[1]",
		},
		{
			Name:     "pattern 4",
			Arr1:     "[]",
			Arr2:     "[]",
			Expected: "[]",
		},
		{
			Name:     "pattern 5",
			Arr1:     "",
			Arr2:     "[]",
			Expected: "[]",
		},
		{
			Name:     "pattern 6",
			Arr1:     "",
			Arr2:     "[1]",
			Expected: "[1]",
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			v, err := concatJSONArray(data.Arr1, data.Arr2)
			if err != nil {
				t.Fatalf("%+v\n", err)
			}
			if v != data.Expected {
				t.Errorf("result of concatJSONArray() is unmatched with expected.\nArr1: %v\nArr2: %v\nExpected: %#08x\nActual:   %#08x", data.Arr1, data.Arr2, data.Expected, v)
			}
		})
	}
}
