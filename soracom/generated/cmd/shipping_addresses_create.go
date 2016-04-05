package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var ShippingAddressesCreateCmdAddressLine1 string

var ShippingAddressesCreateCmdAddressLine2 string

var ShippingAddressesCreateCmdBuilding string

var ShippingAddressesCreateCmdCity string

var ShippingAddressesCreateCmdCompanyName string

var ShippingAddressesCreateCmdDepartment string

var ShippingAddressesCreateCmdFullName string

var ShippingAddressesCreateCmdOperatorId string

var ShippingAddressesCreateCmdPhoneNumber string

var ShippingAddressesCreateCmdState string

var ShippingAddressesCreateCmdZipCode string




var ShippingAddressesCreateCmdBody string


func init() {
  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdAddressLine1, "address-line1", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdAddressLine2, "address-line2", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdBuilding, "building", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdCity, "city", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdCompanyName, "company-name", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdDepartment, "department", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdFullName, "full-name", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdOperatorId, "operator-id", "", "Operator ID")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdPhoneNumber, "phone-number", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdState, "state", "", "")

  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdZipCode, "zip-code", "", "")



  ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  ShippingAddressesCmd.AddCommand(ShippingAddressesCreateCmd)
}

var ShippingAddressesCreateCmd = &cobra.Command{
  Use: "create",
  Short: TR("Create shipping address"),
  Long: TR(`商品発送先を新規登録する。`),
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
    
    param, err := collectShippingAddressesCreateCmdParams()
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

func collectShippingAddressesCreateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForShippingAddressesCreateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForShippingAddressesCreateCmd("/operators/{operator_id}/shipping_addresses"),
    query: buildQueryForShippingAddressesCreateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForShippingAddressesCreateCmd(path string) string {
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", ShippingAddressesCreateCmdOperatorId, -1)
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForShippingAddressesCreateCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForShippingAddressesCreateCmd() (string, error) {
  if ShippingAddressesCreateCmdBody != "" {
    if strings.HasPrefix(ShippingAddressesCreateCmdBody, "@") {
      fname := strings.TrimPrefix(ShippingAddressesCreateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if ShippingAddressesCreateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return ShippingAddressesCreateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if ShippingAddressesCreateCmdAddressLine1 != "" {
    result["addressLine1"] = ShippingAddressesCreateCmdAddressLine1
  }
  
  if ShippingAddressesCreateCmdAddressLine2 != "" {
    result["addressLine2"] = ShippingAddressesCreateCmdAddressLine2
  }
  
  if ShippingAddressesCreateCmdBuilding != "" {
    result["building"] = ShippingAddressesCreateCmdBuilding
  }
  
  if ShippingAddressesCreateCmdCity != "" {
    result["city"] = ShippingAddressesCreateCmdCity
  }
  
  if ShippingAddressesCreateCmdCompanyName != "" {
    result["companyName"] = ShippingAddressesCreateCmdCompanyName
  }
  
  if ShippingAddressesCreateCmdDepartment != "" {
    result["department"] = ShippingAddressesCreateCmdDepartment
  }
  
  if ShippingAddressesCreateCmdFullName != "" {
    result["fullName"] = ShippingAddressesCreateCmdFullName
  }
  
  if ShippingAddressesCreateCmdOperatorId != "" {
    result["operator_id"] = ShippingAddressesCreateCmdOperatorId
  }
  
  if ShippingAddressesCreateCmdPhoneNumber != "" {
    result["phoneNumber"] = ShippingAddressesCreateCmdPhoneNumber
  }
  
  if ShippingAddressesCreateCmdState != "" {
    result["state"] = ShippingAddressesCreateCmdState
  }
  
  if ShippingAddressesCreateCmdZipCode != "" {
    result["zipCode"] = ShippingAddressesCreateCmdZipCode
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

