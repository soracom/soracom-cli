package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// jsonErrorsEnabled reports whether structured JSON error output is opted in
// via the SORACOM_JSON_ERRORS environment variable. It is off by default so the
// CLI keeps its conventional cobra error/usage behavior; agents that want
// machine-readable failures set the variable to any non-empty value.
func jsonErrorsEnabled() bool {
	return os.Getenv("SORACOM_JSON_ERRORS") != ""
}

// PrintError writes the given error to stderr as structured JSON so that AI
// agents and scripts can parse failures without heuristics. API errors already
// carry a JSON body from the SORACOM API and are emitted as-is; any other error
// is wrapped so that the output is always valid JSON.
//
// This only takes effect when SORACOM_JSON_ERRORS is set; otherwise cobra has
// already printed the error in its default format and this is a no-op.
func PrintError(err error) {
	if !jsonErrorsEnabled() {
		return
	}
	printErrorTo(os.Stderr, err)
}

func printErrorTo(w io.Writer, err error) {
	if err == nil {
		return
	}

	var ae *apiError
	if errors.As(err, &ae) {
		body := strings.TrimSpace(ae.ResponseBody)
		if json.Valid([]byte(body)) {
			fmt.Fprintln(w, body)
			return
		}
		printWrappedErrorTo(w, body)
		return
	}

	printWrappedErrorTo(w, err.Error())
}

func printWrappedErrorTo(w io.Writer, message string) {
	out := map[string]interface{}{
		"error": map[string]interface{}{
			"message": message,
		},
	}
	b, err := json.Marshal(out)
	if err != nil {
		fmt.Fprintln(w, message)
		return
	}
	fmt.Fprintln(w, string(b))
}
