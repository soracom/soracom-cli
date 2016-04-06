package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var VpgGetCmdVpgId string





func init() {
  VpgGetCmd.Flags().StringVar(&VpgGetCmdVpgId, "vpg-id", "", "対象のVPGのID")




  VpgCmd.AddCommand(VpgGetCmd)
}

var VpgGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get Virtual Private Gateway"),
  Long: TR(`指定されたVPGの情報を取得する`),
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
    
    param, err := collectVpgGetCmdParams()
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

func collectVpgGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForVpgGetCmd("/virtual_private_gateways/{vpg_id}"),
    query: buildQueryForVpgGetCmd(),
    
    
  }, nil
}

func buildPathForVpgGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "vpg_id" + "}", VpgGetCmdVpgId, -1)
  
  
  
  
  return path
}

func buildQueryForVpgGetCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


