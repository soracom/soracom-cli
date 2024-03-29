package cmd

import (
	"net/http"
	"os"
	"testing"
	"time"

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
	time.Sleep(1 * time.Second)
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
			assert.NoError(t, err)
			assert.Equal(t, data.Expected, v)
		})
	}
}

func TestHideSecretHeaders(t *testing.T) {
	var testData = []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name:     "Nominal case for secret headers: X-Soracom-Api-Key",
			Input:    "GET / HTTP/1.1\nX-Soracom-Api-Key: this-should-be-hidden\nHello: world",
			Expected: "GET / HTTP/1.1\nX-Soracom-Api-Key: <hidden>\nHello: world",
		},
		{
			Name:     "Nominal case for secret headers: X-Soracom-Token",
			Input:    "GET / HTTP/1.1\nX-Soracom-Token: this-should-be-hidden\nHello: world",
			Expected: "GET / HTTP/1.1\nX-Soracom-Token: <hidden>\nHello: world",
		},
		{
			Name:     "Nominal case for headers",
			Input:    "GET / HTTP/1.1\nHeader: leave-this-as-it-is\nHello: world",
			Expected: "GET / HTTP/1.1\nHeader: leave-this-as-it-is\nHello: world",
		},
		{
			Name:     "Headers which contain secret header but are not exactly same with secret header should not be replaced",
			Input:    "GET / HTTP/1.1\nX-Soracom-Api-Key-Version: 2022-04-22\nHello: world",
			Expected: "GET / HTTP/1.1\nX-Soracom-Api-Key-Version: 2022-04-22\nHello: world",
		},
		{
			Name:     "Headers should be treated as case-insensitive and transparently",
			Input:    "GET / HTTP/1.1\nx-sORACOM-aPI-kEY: this-should-be-hidden\nHello: world",
			Expected: "GET / HTTP/1.1\nx-sORACOM-aPI-kEY: <hidden>\nHello: world",
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			actualByte := hideSecretHeaders([]byte(data.Input))
			actual := string(actualByte)
			if actual != data.Expected {
				t.Errorf("result of hideSecureHeaders() is unmatched with expected.\nExpected: %v\nActual: %v", data.Expected, actual)
			}
		})
	}
}
