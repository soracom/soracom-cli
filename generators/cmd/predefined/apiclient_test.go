package cmd

import (
	"fmt"
	"net/http"
	"os"
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
	origEnvVars := saveEnvVars([]string{"HTTP_PROXY"})
	defer restoreEnvVars(origEnvVars)

	proxyAddr := "localhost:18080"
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

	ac2 := newAPIClient(&apiClientOptions{})
	ac2.callAPI(&apiParams{
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
