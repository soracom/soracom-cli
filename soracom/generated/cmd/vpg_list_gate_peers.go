package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var VpgListGatePeersCmdVpgId string






func init() {
  VpgListGatePeersCmd.Flags().StringVar(&VpgListGatePeersCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.list_virtual_private_gateway_peers.get.parameters.vpg_id.description"))




  VpgCmd.AddCommand(VpgListGatePeersCmd)
}

var VpgListGatePeersCmd = &cobra.Command{
  Use: "list-gate-peers",
  Short: TR("virtual_private_gateway.list_virtual_private_gateway_peers.get.summary"),
  Long: TR(`virtual_private_gateway.list_virtual_private_gateway_peers.get.description`),
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
    
    param, err := collectVpgListGatePeersCmdParams()
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

func collectVpgListGatePeersCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForVpgListGatePeersCmd("/virtual_private_gateways/{vpg_id}/gate/peers"),
    query: buildQueryForVpgListGatePeersCmd(),
    
    
  }, nil
}

func buildPathForVpgListGatePeersCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "vpg_id" + "}", VpgListGatePeersCmdVpgId, -1)
  
  
  
  
  
  return path
}

func buildQueryForVpgListGatePeersCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


