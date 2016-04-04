package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var ShippingAddressesListCmdOperatorId string





func init() {
  ShippingAddressesListCmd.Flags().StringVar(&ShippingAddressesListCmdOperatorId, "operator-id", "", "Operator ID")




  ShippingAddressesCmd.AddCommand(ShippingAddressesListCmd)
}

var ShippingAddressesListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List shipping addresses"),
  Long: TR(`商品発送先の一覧を返します。`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
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
    
    param, err := collectShippingAddressesListCmdParams()
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

func collectShippingAddressesListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForShippingAddressesListCmd("/operators/{operator_id}/shipping_addresses"),
    query: buildQueryForShippingAddressesListCmd(),
    
    
  }, nil
}

func buildPathForShippingAddressesListCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", ShippingAddressesListCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForShippingAddressesListCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


