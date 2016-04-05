package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersListCmdOperatorId string





func init() {
  UsersListCmd.Flags().StringVar(&UsersListCmdOperatorId, "operator-id", "", "operator_id")




  UsersCmd.AddCommand(UsersListCmd)
}

var UsersListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List Users"),
  Long: TR(`SAMユーザー一覧を返す。`),
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
    
    param, err := collectUsersListCmdParams()
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

func collectUsersListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForUsersListCmd("/operators/{operator_id}/users"),
    query: buildQueryForUsersListCmd(),
    
    
  }, nil
}

func buildPathForUsersListCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersListCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForUsersListCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


