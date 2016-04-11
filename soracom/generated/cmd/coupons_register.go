package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var CouponsRegisterCmdCouponCode string






func init() {
  CouponsRegisterCmd.Flags().StringVar(&CouponsRegisterCmdCouponCode, "coupon-code", "", "coupon_code")




  CouponsCmd.AddCommand(CouponsRegisterCmd)
}

var CouponsRegisterCmd = &cobra.Command{
  Use: "register",
  Short: TR("Register Coupon"),
  Long: TR(`クーポンを登録します。`),
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
    
    param, err := collectCouponsRegisterCmdParams()
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

func collectCouponsRegisterCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForCouponsRegisterCmd("/coupons/{coupon_code}/register"),
    query: buildQueryForCouponsRegisterCmd(),
    
    
  }, nil
}

func buildPathForCouponsRegisterCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "coupon_code" + "}", CouponsRegisterCmdCouponCode, -1)
  
  
  
  
  
  return path
}

func buildQueryForCouponsRegisterCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


