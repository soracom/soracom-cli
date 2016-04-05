package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OrdersCreateCmdShippingAddressId string




var OrdersCreateCmdBody string


func init() {
  OrdersCreateCmd.Flags().StringVar(&OrdersCreateCmdShippingAddressId, "shipping-address-id", "", "")



  OrdersCreateCmd.Flags().StringVar(&OrdersCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  OrdersCmd.AddCommand(OrdersCreateCmd)
}

var OrdersCreateCmd = &cobra.Command{
  Use: "create",
  Short: TR("Create Quotation"),
  Long: TR(`新規で発注見積もりの作成を行います。orderIdを/confirmにPUTすると発注が完了します。`),
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
    
    param, err := collectOrdersCreateCmdParams()
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

func collectOrdersCreateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForOrdersCreateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForOrdersCreateCmd("/orders"),
    query: buildQueryForOrdersCreateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForOrdersCreateCmd(path string) string {
  
  
  
  
  
  return path
}

func buildQueryForOrdersCreateCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForOrdersCreateCmd() (string, error) {
  if OrdersCreateCmdBody != "" {
    if strings.HasPrefix(OrdersCreateCmdBody, "@") {
      fname := strings.TrimPrefix(OrdersCreateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if OrdersCreateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return OrdersCreateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if OrdersCreateCmdShippingAddressId != "" {
    result["shippingAddressId"] = OrdersCreateCmdShippingAddressId
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

