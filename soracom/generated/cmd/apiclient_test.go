package cmd

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"testing"

	"github.com/elazarl/goproxy"
)

type dummyAPIServer struct {
	accessCount int
}

func (s *dummyAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.accessCount++
	fmt.Fprintf(w, "response from dummy api server")
}

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func TestCallAPIWithProxy(t *testing.T) {
	if runtime.GOOS != "darwin" {
		// we are not able to get goproxy working on Linux env
		return
	}
	origEnvVars := saveEnvVars([]string{"HTTP_PROXY"})
	defer restoreEnvVars(origEnvVars)

	proxyAddr := ":18080"
	err := os.Setenv("HTTP_PROXY", proxyAddr)
	if err != nil {
		t.Fatalf("os.Setenv() failed.")
	}

	var proxyAccessCount int
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	var CountAndReject goproxy.FuncHttpsHandler = func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		proxyAccessCount++
		return goproxy.RejectConnect, host
	}

	proxy.OnRequest().HandleConnect(CountAndReject)
	go http.ListenAndServe(proxyAddr, proxy)

	ac := newAPIClient(&apiClientOptions{})
	ac.callAPI(&apiParams{
		method:         "GET",
		path:           "v1/subscribers",
		noRetryOnError: true,
	})
	if proxyAccessCount == 0 {
		t.Fatalf("proxy server should be accessed when HTTP_PROXY env var is set")
	}
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

func TestVersionInt(t *testing.T) {
	var testData = []struct {
		Name   string
		VerStr string
		VerInt uint32
	}{
		{
			Name:   "pattern 1",
			VerStr: "v0.1.2",
			VerInt: 0x00010200,
		},
		{
			Name:   "pattern 2",
			VerStr: "v1.22.333", // 333 == 0x14d but only 0x4d will be stored in the 3rd place
			VerInt: 0x01164d00,
		},
		{
			Name:   "pattern 3",
			VerStr: "1.2.3", // no "v" prefix
			VerInt: 0x01020300,
		},
		{
			Name:   "pattern 4",
			VerStr: "v1.2.3-special",
			VerInt: 0x01020300,
		},
		{
			Name:   "pattern 5",
			VerStr: "v1.2.3-special1",
			VerInt: 0x01020301,
		},
		{
			Name:   "pattern 6",
			VerStr: "v1.2.3.4",
			VerInt: 0x01020304,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			v := versionInt(data.VerStr)
			if v != data.VerInt {
				t.Errorf("result of versionInt() is unmatched with expected.\nArg: %v\nExpected: %#08x\nActual:   %#08x", data.VerStr, data.VerInt, v)
			}
		})
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
