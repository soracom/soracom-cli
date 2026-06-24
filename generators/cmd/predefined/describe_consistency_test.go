package cmd

import (
	"strings"
	"sync"
	"testing"

	"github.com/spf13/pflag"
)

var initCommandsOnce sync.Once

func ensureCommandsInitialized() {
	initCommandsOnce.Do(InitRootCmd)
}

// TestDescribeOptionsMatchGeneratedFlags is the safety net for the rule
// duplication between `describe` and the code generator: `describe` re-derives
// option names, flag eligibility and the operator_id special case from the
// OpenAPI definition, while the actual flags are produced by the generator. For
// every command `describe` knows about, the set of flag options it reports must
// equal the flags actually registered on the corresponding cobra command
// (excluding global and common flags). If the two ever disagree, this test
// fails instead of shipping a misleading schema.
func TestDescribeOptionsMatchGeneratedFlags(t *testing.T) {
	ensureCommandsInitialized()

	entries, err := loadDescribeEntries()
	if err != nil {
		t.Fatalf("loadDescribeEntries: %v", err)
	}

	common := map[string]bool{"body": true, "fetch-all": true, "jsonl": true, "help": true}

	for _, e := range entries {
		c, _, err := RootCmd.Find(strings.Fields(e.command))
		if err != nil || c == nil {
			t.Errorf("%s: no matching cobra command found", e.command)
			continue
		}

		// LocalFlags() are the command's own flags (param/body flags); the
		// inherited global flags are excluded automatically.
		flagOpts := map[string]bool{}
		c.LocalFlags().VisitAll(func(f *pflag.Flag) {
			if !common[f.Name] {
				flagOpts[f.Name] = true
			}
		})

		descOpts := map[string]bool{}
		d := buildCommandDescription(e)
		for _, p := range d.Parameters {
			if p.Option != "" {
				descOpts[p.Option] = true
			}
		}
		if d.RequestBody != nil {
			for _, p := range d.RequestBody.Properties {
				if p.Option != "" {
					descOpts[p.Option] = true
				}
			}
		}

		for opt := range descOpts {
			if !flagOpts[opt] {
				t.Errorf("%s: describe reports option --%s with no matching command flag", e.command, opt)
			}
		}
		for fl := range flagOpts {
			if !descOpts[fl] {
				t.Errorf("%s: command flag --%s is not reported by describe", e.command, fl)
			}
		}
	}
}
