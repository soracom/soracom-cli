package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersGetCmdOperatorId string

var UsersGetCmdUserName string





func init() {
  UsersGetCmd.Flags().StringVar(&UsersGetCmdOperatorId, "operator-id", "", "operator_id")

  UsersGetCmd.Flags().StringVar(&UsersGetCmdUserName, "user-name", "", "user_name")




  UsersCmd.AddCommand(UsersGetCmd)
}

var UsersGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get User"),
  Long: TR(`SAMユーザーを返す。`),
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
    
    param, err := collectUsersGetCmdParams()
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

func collectUsersGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForUsersGetCmd("/operators/{operator_id}/users/{user_name}"),
    query: buildQueryForUsersGetCmd(),
    
    
  }, nil
}

func buildPathForUsersGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersGetCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersGetCmdUserName, -1)
  
  
  
  
  return path
}

func buildQueryForUsersGetCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


