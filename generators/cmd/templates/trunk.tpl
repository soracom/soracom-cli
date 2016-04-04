package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  {{.ParentCommandVariableName}}.AddCommand({{.CommandVariableName}})
}

var {{.CommandVariableName}} = &cobra.Command{
  Use: "{{.Use}}",
  Short: TR("{{.Short}}"),
  Long: TR(`{{.Long}}`),
}
