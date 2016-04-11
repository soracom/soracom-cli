package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var ShippingAddressesUpdateCmdAddressLine1 string

var ShippingAddressesUpdateCmdAddressLine2 string

var ShippingAddressesUpdateCmdBuilding string

var ShippingAddressesUpdateCmdCity string

var ShippingAddressesUpdateCmdCompanyName string

var ShippingAddressesUpdateCmdDepartment string

var ShippingAddressesUpdateCmdFullName string

var ShippingAddressesUpdateCmdOperatorId string

var ShippingAddressesUpdateCmdPhoneNumber string

var ShippingAddressesUpdateCmdShippingAddressId string

var ShippingAddressesUpdateCmdState string

var ShippingAddressesUpdateCmdZipCode string





var ShippingAddressesUpdateCmdBody string


func init() {
  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdAddressLine1, "address-line1", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdAddressLine2, "address-line2", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdBuilding, "building", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdCity, "city", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdCompanyName, "company-name", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdDepartment, "department", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdFullName, "full-name", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdOperatorId, "operator-id", "", "Operator ID")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdPhoneNumber, "phone-number", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdShippingAddressId, "shipping-address-id", "", "shipping_address_id")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdState, "state", "", "")

  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdZipCode, "zip-code", "", "")



  ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  ShippingAddressesCmd.AddCommand(ShippingAddressesUpdateCmd)
}

var ShippingAddressesUpdateCmd = &cobra.Command{
  Use: "update",
  Short: TR("Update shipping address"),
  Long: TR(`商品発送先を更新する。`),
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
    
    param, err := collectShippingAddressesUpdateCmdParams()
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

func collectShippingAddressesUpdateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForShippingAddressesUpdateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "PUT",
    path: buildPathForShippingAddressesUpdateCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
    query: buildQueryForShippingAddressesUpdateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForShippingAddressesUpdateCmd(path string) string {
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", ShippingAddressesUpdateCmdOperatorId, -1)
  
  
  
  
  
  path = strings.Replace(path, "{" + "shipping_address_id" + "}", ShippingAddressesUpdateCmdShippingAddressId, -1)
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForShippingAddressesUpdateCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForShippingAddressesUpdateCmd() (string, error) {
  if ShippingAddressesUpdateCmdBody != "" {
    if strings.HasPrefix(ShippingAddressesUpdateCmdBody, "@") {
      fname := strings.TrimPrefix(ShippingAddressesUpdateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if ShippingAddressesUpdateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return ShippingAddressesUpdateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if ShippingAddressesUpdateCmdAddressLine1 != "" {
    result["addressLine1"] = ShippingAddressesUpdateCmdAddressLine1
  }
  
  
  
  if ShippingAddressesUpdateCmdAddressLine2 != "" {
    result["addressLine2"] = ShippingAddressesUpdateCmdAddressLine2
  }
  
  
  
  if ShippingAddressesUpdateCmdBuilding != "" {
    result["building"] = ShippingAddressesUpdateCmdBuilding
  }
  
  
  
  if ShippingAddressesUpdateCmdCity != "" {
    result["city"] = ShippingAddressesUpdateCmdCity
  }
  
  
  
  if ShippingAddressesUpdateCmdCompanyName != "" {
    result["companyName"] = ShippingAddressesUpdateCmdCompanyName
  }
  
  
  
  if ShippingAddressesUpdateCmdDepartment != "" {
    result["department"] = ShippingAddressesUpdateCmdDepartment
  }
  
  
  
  if ShippingAddressesUpdateCmdFullName != "" {
    result["fullName"] = ShippingAddressesUpdateCmdFullName
  }
  
  
  
  
  
  if ShippingAddressesUpdateCmdPhoneNumber != "" {
    result["phoneNumber"] = ShippingAddressesUpdateCmdPhoneNumber
  }
  
  
  
  
  
  if ShippingAddressesUpdateCmdState != "" {
    result["state"] = ShippingAddressesUpdateCmdState
  }
  
  
  
  if ShippingAddressesUpdateCmdZipCode != "" {
    result["zipCode"] = ShippingAddressesUpdateCmdZipCode
  }
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

