package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var ShippingAddressesDeleteCmdOperatorId string

var ShippingAddressesDeleteCmdShippingAddressId string





func init() {
  ShippingAddressesDeleteCmd.Flags().StringVar(&ShippingAddressesDeleteCmdOperatorId, "operator-id", "", "Operator ID")

  ShippingAddressesDeleteCmd.Flags().StringVar(&ShippingAddressesDeleteCmdShippingAddressId, "shipping-address-id", "", "shipping_address_id")




  ShippingAddressesCmd.AddCommand(ShippingAddressesDeleteCmd)
}

var ShippingAddressesDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("Delete shipping address"),
  Long: TR(`商品発送先を削除する。`),
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
    
    param, err := collectShippingAddressesDeleteCmdParams()
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

func collectShippingAddressesDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForShippingAddressesDeleteCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
    query: buildQueryForShippingAddressesDeleteCmd(),
    
    
  }, nil
}

func buildPathForShippingAddressesDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", ShippingAddressesDeleteCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "shipping_address_id" + "}", ShippingAddressesDeleteCmdShippingAddressId, -1)
  
  
  
  
  return path
}

func buildQueryForShippingAddressesDeleteCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


