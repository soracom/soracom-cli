package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var ShippingAddressesGetCmdOperatorId string

var ShippingAddressesGetCmdShippingAddressId string





func init() {
  ShippingAddressesGetCmd.Flags().StringVar(&ShippingAddressesGetCmdOperatorId, "operator-id", "", "Operator ID")

  ShippingAddressesGetCmd.Flags().StringVar(&ShippingAddressesGetCmdShippingAddressId, "shipping-address-id", "", "shipping_address_id")




  ShippingAddressesCmd.AddCommand(ShippingAddressesGetCmd)
}

var ShippingAddressesGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get shipping address"),
  Long: TR(`商品発送先を返す。`),
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
    
    param, err := collectShippingAddressesGetCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    return prettyPrintJSON(result)
  },
}

func collectShippingAddressesGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForShippingAddressesGetCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
    query: buildQueryForShippingAddressesGetCmd(),
    
    
  }, nil
}

func buildPathForShippingAddressesGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", ShippingAddressesGetCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "shipping_address_id" + "}", ShippingAddressesGetCmdShippingAddressId, -1)
  
  
  
  
  return path
}

func buildQueryForShippingAddressesGetCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


