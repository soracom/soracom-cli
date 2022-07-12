package main

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tj/assert"
)

const apiDef1String = `
paths:
  "/v1/subscribers":
    get:
      x-soracom-cli:
        - subscribers list
  '/v1/subscribers/{imsi}':
    get:
      x-soracom-cli:
        - subscribers get
  '/v1/stats/air':
    get:
      x-soracom-cli:
        - stats air
  '/v1/stats/beam/hoge':
    get:
      x-soracom-cli:
        - stats beam hoge
  '/v1/stats/napter/audit-logs':
    get:
      x-soracom-cli:
        - stats napter audit-logs get
`

func TestGenerateArgsForTrunkCommands(t *testing.T) {
	apiDef1, err := openapi3.NewLoader().LoadFromData([]byte(apiDef1String))
	assert.NoError(t, err)

	var testData = []struct {
		Name        string
		APIDef      *openapi3.T
		CommandArgs []commandArgs
	}{
		{
			Name:   "pattern 1",
			APIDef: apiDef1,
			CommandArgs: []commandArgs{
				{
					Use:                       "stats",
					Short:                     "cli.stats.summary",
					Long:                      "cli.stats.description",
					CommandVariableName:       "StatsCmd",
					ParentCommandVariableName: "RootCmd",
					FileName:                  "stats.go",
				},
				{
					Use:                       "beam",
					Short:                     "cli.stats.beam.summary",
					Long:                      "cli.stats.beam.description",
					CommandVariableName:       "StatsBeamCmd",
					ParentCommandVariableName: "StatsCmd",
					FileName:                  "stats_beam.go",
				},
				{
					Use:                       "napter",
					Short:                     "cli.stats.napter.summary",
					Long:                      "cli.stats.napter.description",
					CommandVariableName:       "StatsNapterCmd",
					ParentCommandVariableName: "StatsCmd",
					FileName:                  "stats_napter.go",
				},
				{
					Use:                       "audit-logs",
					Short:                     "cli.stats.napter.audit-logs.summary",
					Long:                      "cli.stats.napter.audit-logs.description",
					CommandVariableName:       "StatsNapterAuditLogsCmd",
					ParentCommandVariableName: "StatsNapterCmd",
					FileName:                  "stats_napter_audit_logs.go",
				},
				{
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
	apiDef1, err := openapi3.NewLoader().LoadFromData([]byte(apiDef1String))
	assert.NoError(t, err)

	var testData = []struct {
		Name          string
		APIDef        *openapi3.T
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
