package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var VpgTerminateCmdVpgId string






func init() {
  VpgTerminateCmd.Flags().StringVar(&VpgTerminateCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.terminate_virtual_private_gateway.post.parameters.vpg_id.description"))




  VpgCmd.AddCommand(VpgTerminateCmd)
}

var VpgTerminateCmd = &cobra.Command{
  Use: "terminate",
  Short: TR("virtual_private_gateway.terminate_virtual_private_gateway.post.summary"),
  Long: TR(`virtual_private_gateway.terminate_virtual_private_gateway.post.description`),
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
    
    param, err := collectVpgTerminateCmdParams()
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

func collectVpgTerminateCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForVpgTerminateCmd("/virtual_private_gateways/{vpg_id}/terminate"),
    query: buildQueryForVpgTerminateCmd(),
    
    
  }, nil
}

func buildPathForVpgTerminateCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "vpg_id" + "}", VpgTerminateCmdVpgId, -1)
  
  
  
  
  
  return path
}

func buildQueryForVpgTerminateCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


