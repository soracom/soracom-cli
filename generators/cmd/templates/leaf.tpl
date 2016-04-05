package cmd

import (
{{if .BodyExists }}
  "encoding/json"
{{end}}
  "fmt"
{{if .BodyExists }}
  "io/ioutil"
{{end}}
  "os"
  "strings"

  "github.com/spf13/cobra"
)

{{$prefix := .CommandVariableName}}
{{$suffix := .CommandVariableName}}
{{$cmdvar := .CommandVariableName}}

{{range .StringFlags}}
var {{$prefix}}{{.VarName}} string
{{end}}
{{range .IntegerFlags}}
var {{$prefix}}{{.VarName}} int64
{{end}}
{{range .FloatFlags}}
var {{$prefix}}{{.VarName}} float64
{{end}}
{{if .BodyExists }}
var {{$prefix}}Body string
{{end}}

func init() {
{{- range .StringFlags}}
  {{$cmdvar}}.Flags().StringVar(&{{$prefix}}{{.VarName}}, "{{.LongOption}}", "{{.DefaultValue}}", "{{.ShortHelp}}")
{{end}}

{{- range .IntegerFlags}}
  {{$cmdvar}}.Flags().Int64Var(&{{$prefix}}{{.VarName}}, "{{.LongOption}}", {{.DefaultValue}}, "{{.ShortHelp}}")
{{end}}

{{- range .FloatFlags}}
  {{$cmdvar}}.Flags().Float64Var(&{{$prefix}}{{.VarName}}, "{{.LongOption}}", {{.DefaultValue}}, "{{.ShortHelp}}")
{{end}}

{{if .BodyExists }}
  {{$cmdvar}}.Flags().StringVar(&{{$prefix}}Body, "body", "", TR("cli.common_params.body.short_help"))
{{end}}

  {{.ParentCommandVariableName}}.AddCommand({{$cmdvar}})
}

var {{ $cmdvar }} = &cobra.Command{
  Use: "{{.Use}}",
  Short: TR("{{.Short}}"),
  Long: TR(`{{.Long}}`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "{{.BasePath}}",
    }

    ac := newAPIClient(opt)
    if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
      ac.SetVerbose(true)
    }

    {{if .RequireAuth}}
    err := authHelper(ac, cmd, args)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }
    {{end}}
    param, err := collect{{$cmdvar}}Params()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    fmt.Println(result)
    return nil
  },
}

func collect{{$cmdvar}}Params() (*apiParams, error) {
  {{if .BodyExists }}
  body, err := buildBodyFor{{$suffix}}()
  if err != nil {
    return nil, err
  }
  {{end}}

  return &apiParams{
    method: "{{.Method}}",
    path: buildPathFor{{$suffix}}("{{.Path}}"),
    query: buildQueryFor{{$suffix}}(),
    {{if .BodyExists }}contentType: "{{.ContentType}}",{{end}}
    {{if .BodyExists }}body: body,{{end}}
  }, nil
}

func buildPathFor{{$suffix}}(path string) string {
  {{range .StringFlags}}
  {{if eq .In "path"}}
  path = strings.Replace(path, "{" + "{{.Name}}" + "}", {{$prefix}}{{.VarName}}, -1)
  {{end}}
  {{end}}
  {{range .IntegerFlags}}
  {{if eq .In "path"}}
  path = strings.Replace(path, "{" + "{{.Name}}" + "}", {{$prefix}}{{.VarName}}, -1)
  {{end}}
  {{end}}
  {{range .FloatFlags}}
  {{if eq .In "path"}}
  path = strings.Replace(path, "{" + "{{.Name}}" + "}", {{$prefix}}{{.VarName}}, -1)
  {{end}}
  {{end}}
  return path
}

func buildQueryFor{{$suffix}}() string {
  result := []string{}
  {{range .StringFlags}}
  {{if eq .In "query"}}
  if {{$prefix}}{{.VarName}} != "{{.DefaultValue}}" {
    result = append(result, fmt.Sprintf("%s=%s", "{{.Name}}", {{$prefix}}{{.VarName}}))
  }
  {{end}}
  {{end}}

  {{range .IntegerFlags}}
  {{if eq .In "query"}}
  if {{$prefix}}{{.VarName}} != {{.DefaultValue}} {
    result = append(result, fmt.Sprintf("%s=%d", "{{.Name}}", {{$prefix}}{{.VarName}}))
  }
  {{end}}
  {{end}}

  {{range .FloatFlags}}
  {{if eq .In "query"}}
  if {{$prefix}}{{.VarName}} != {{.DefaultValue}} {
    result = append(result, fmt.Sprintf("%s=%f", "{{.Name}}", {{$prefix}}{{.VarName}}))
  }
  {{end}}
  {{end}}
  return strings.Join(result, "&")
}

{{if .BodyExists }}
func buildBodyFor{{$suffix}}() (string, error) {
  if {{$prefix}}Body != "" {
    if strings.HasPrefix({{$prefix}}Body, "@") {
      fname := strings.TrimPrefix({{$prefix}}Body, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if {{$prefix}}Body == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return {{$prefix}}Body, nil
    }
  }

  result := map[string]interface{}{}
  {{range .StringFlags}}
  if {{$prefix}}{{.VarName}} != "{{.DefaultValue}}" {
    result["{{.Name}}"] = {{$prefix}}{{.VarName}}
  }
  {{end}}
  {{range .IntegerFlags}}
  if {{$prefix}}{{.VarName}} != {{.DefaultValue}} {
    result["{{.Name}}"] = {{$prefix}}{{.VarName}}
  }
  {{end}}
  {{range .FloatFlags}}
  if {{$prefix}}{{.VarName}} != {{.DefaultValue}} {
    result["{{.Name}}"] = {{$prefix}}{{.VarName}}
  }
  {{end}}

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}
{{end}}