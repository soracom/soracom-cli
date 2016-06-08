package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersPasswordDeleteCmdOperatorId string

var UsersPasswordDeleteCmdUserName string






func init() {
  UsersPasswordDeleteCmd.Flags().StringVar(&UsersPasswordDeleteCmdOperatorId, "operator-id", "", TR("operator_id"))

  UsersPasswordDeleteCmd.Flags().StringVar(&UsersPasswordDeleteCmdUserName, "user-name", "", TR("user_name"))




  UsersPasswordCmd.AddCommand(UsersPasswordDeleteCmd)
}

var UsersPasswordDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("users.delete_user_password.delete.summary"),
  Long: TR(`users.delete_user_password.delete.description`),
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
    
    param, err := collectUsersPasswordDeleteCmdParams()
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

func collectUsersPasswordDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForUsersPasswordDeleteCmd("/operators/{operator_id}/users/{user_name}/password"),
    query: buildQueryForUsersPasswordDeleteCmd(),
    
    
  }, nil
}

func buildPathForUsersPasswordDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersPasswordDeleteCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersPasswordDeleteCmdUserName, -1)
  
  
  
  
  
  return path
}

func buildQueryForUsersPasswordDeleteCmd() string {
  result := []string{}
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


