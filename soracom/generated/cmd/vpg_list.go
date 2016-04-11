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
  VpgListCmd.Flags().StringVar(&VpgListCmdLastEvaluatedKey, "last-evaluated-key", "", "現ページで取得した最後のGroupのID。このパラメータを指定することで次のVPGから始まるリストを取得できる。")

  VpgListCmd.Flags().StringVar(&VpgListCmdTagName, "tag-name", "", "VPGについたTag名。完全一致するTag名が検索対象となる。tag_nameを指定した場合、tag_valueはが必須。")

  VpgListCmd.Flags().StringVar(&VpgListCmdTagValue, "tag-value", "", "VPGについたTagの値。")

  VpgListCmd.Flags().StringVar(&VpgListCmdTagValueMatchMode, "tag-value-match-mode", "", "タグの検索条件。")

  VpgListCmd.Flags().Int64Var(&VpgListCmdLimit, "limit", 0, "レスポンス1ページあたりの最大数")




  VpgCmd.AddCommand(VpgListCmd)
}

var VpgListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List Virtual Private Gateways"),
  Long: TR(`VPGの一覧を返す
`),
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


