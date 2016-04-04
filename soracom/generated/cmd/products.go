package cmd

import (
  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(ProductsCmd)
}

var ProductsCmd = &cobra.Command{
  Use: "products",
  Short: TR("products.cli.summary"),
  Long: TR(`products.cli.description`),
}
