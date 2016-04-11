package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)











func init() {



  ProductsCmd.AddCommand(ProductsListCmd)
}

var ProductsListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List products"),
  Long: TR(`商品一覧を返します。`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
      Language: getSelectedLanguage(),
    }

    ac := newAPIClient(opt)
    if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
      ac.SetVerbose(true)
    }

    
    err := authHelper(ac, cmd, args)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }
    
    param, err := collectProductsListCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    if result != "" {
      return prettyPrintStringAsJSON(result)
    } else {
      return nil
    }
  },
}

func collectProductsListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForProductsListCmd("/products"),
    query: buildQueryForProductsListCmd(),
    
    
  }, nil
}

func buildPathForProductsListCmd(path string) string {
  
  
  
  
  return path
}

func buildQueryForProductsListCmd() string {
  result := []string{}
  

  

  

  

  return strings.Join(result, "&")
}


