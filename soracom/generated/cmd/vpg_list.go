package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var VpgListCmdLastEvaluatedKey string

var VpgListCmdTagName string

var VpgListCmdTagValue string

var VpgListCmdTagValueMatchMode string


var VpgListCmdLimit int64





func init() {
  VpgListCmd.Flags().StringVar(&VpgListCmdLastEvaluatedKey, "last-evaluated-key", "", TR("virtual_private_gateway.list_virtual_private_gateways.get.parameters.last_evaluated_key.description"))

  VpgListCmd.Flags().StringVar(&VpgListCmdTagName, "tag-name", "", TR("virtual_private_gateway.list_virtual_private_gateways.get.parameters.tag_name.description"))

  VpgListCmd.Flags().StringVar(&VpgListCmdTagValue, "tag-value", "", TR("virtual_private_gateway.list_virtual_private_gateways.get.parameters.tag_value.description"))

  VpgListCmd.Flags().StringVar(&VpgListCmdTagValueMatchMode, "tag-value-match-mode", "", TR("virtual_private_gateway.list_virtual_private_gateways.get.parameters.tag_value_match_mode.description"))

  VpgListCmd.Flags().Int64Var(&VpgListCmdLimit, "limit", 0, TR("virtual_private_gateway.list_virtual_private_gateways.get.parameters.limit.description"))




  VpgCmd.AddCommand(VpgListCmd)
}

var VpgListCmd = &cobra.Command{
  Use: "list",
  Short: TR("virtual_private_gateway.list_virtual_private_gateways.get.summary"),
  Long: TR(`virtual_private_gateway.list_virtual_private_gateways.get.description`),
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
    
    param, err := collectVpgListCmdParams()
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

func collectVpgListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForVpgListCmd("/virtual_private_gateways"),
    query: buildQueryForVpgListCmd(),
    
    
  }, nil
}

func buildPathForVpgListCmd(path string) string {
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForVpgListCmd() string {
  result := []string{}
  
  
  if VpgListCmdLastEvaluatedKey != "" {
    result = append(result, sprintf("%s=%s", "last_evaluated_key", VpgListCmdLastEvaluatedKey))
  }
  
  
  
  if VpgListCmdTagName != "" {
    result = append(result, sprintf("%s=%s", "tag_name", VpgListCmdTagName))
  }
  
  
  
  if VpgListCmdTagValue != "" {
    result = append(result, sprintf("%s=%s", "tag_value", VpgListCmdTagValue))
  }
  
  
  
  if VpgListCmdTagValueMatchMode != "" {
    result = append(result, sprintf("%s=%s", "tag_value_match_mode", VpgListCmdTagValueMatchMode))
  }
  
  

  
  
  if VpgListCmdLimit != 0 {
    result = append(result, sprintf("%s=%d", "limit", VpgListCmdLimit))
  }
  
  

  

  

  return strings.Join(result, "&")
}


