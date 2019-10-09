package main

import (
	"reflect"
	"testing"

	"github.com/soracom/soracom-cli/generators/lib"
)

var apiDef1 = &lib.APIDefinitions{
	Methods: []lib.APIMethod{
		lib.APIMethod{
			CLI: []string{"subscribers list"},
		},
		lib.APIMethod{
			CLI: []string{"subscribers get"},
		},
		lib.APIMethod{
			CLI: []string{"stats air"},
		},
		lib.APIMethod{
			CLI: []string{"stats beam hoge"},
		},
		lib.APIMethod{
			CLI: []string{"stats napter audit-logs get"},
		},
	},
}

func TestGenerateArgsForTrunkCommands(t *testing.T) {
	var testData = []struct {
		Name        string
		APIDef      *lib.APIDefinitions
		CommandArgs []commandArgs
	}{
		{
			Name:   "pattern 1",
			APIDef: apiDef1,
			CommandArgs: []commandArgs{
				commandArgs{
					Use:                       "stats",
					Short:                     "cli.stats.summary",
					Long:                      "cli.stats.description",
					CommandVariableName:       "StatsCmd",
					ParentCommandVariableName: "RootCmd",
					FileName:                  "stats.go",
				},
				commandArgs{
					Use:                       "beam",
					Short:                     "cli.stats.beam.summary",
					Long:                      "cli.stats.beam.description",
					CommandVariableName:       "StatsBeamCmd",
					ParentCommandVariableName: "StatsCmd",
					FileName:                  "stats_beam.go",
				},
				commandArgs{
					Use:                       "napter",
					Short:                     "cli.stats.napter.summary",
					Long:                      "cli.stats.napter.description",
					CommandVariableName:       "StatsNapterCmd",
					ParentCommandVariableName: "StatsCmd",
					FileName:                  "stats_napter.go",
				},
				commandArgs{
					Use:                       "audit-logs",
					Short:                     "cli.stats.napter.audit-logs.summary",
					Long:                      "cli.stats.napter.audit-logs.description",
					CommandVariableName:       "StatsNapterAuditLogsCmd",
					ParentCommandVariableName: "StatsNapterCmd",
					FileName:                  "stats_napter_audit_logs.go",
				},
				commandArgs{
					Use:                       "subscribers",
					Short:                     "cli.subscribers.summary",
					Long:                      "cli.subscribers.description",
					CommandVariableName:       "SubscribersCmd",
					ParentCommandVariableName: "RootCmd",
					FileName:                  "subscribers.go",
				},
			},
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			a := generateArgsForTrunkCommands(data.APIDef)
			if !reflect.DeepEqual(a, data.CommandArgs) {
				t.Errorf("result of generateArgsForTrunkCommands() is unmatched with expected.\nExpected: %#v\nActual:   %#v\n", data.CommandArgs, a)
			}
		})
	}
}

func TestExtractTrunkCommands(t *testing.T) {
	var testData = []struct {
		Name          string
		APIDef        *lib.APIDefinitions
		TrunkCommands []string
	}{
		{
			Name:   "pattern 1",
			APIDef: apiDef1,
			TrunkCommands: []string{
				"stats",
				"stats beam",
				"stats napter",
				"stats napter audit-logs",
				"subscribers",
			},
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			t.Parallel()

			a := extractTrunkCommands(data.APIDef)
			if !reflect.DeepEqual(a, data.TrunkCommands) {
				t.Errorf("result of extractTrunkCommands() is unmatched with expected.\nExpected: %#v\nActual:   %#v\n", data.TrunkCommands, a)
			}
		})
	}
}
