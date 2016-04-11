package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersAttachRoleCmdOperatorId string

var UsersAttachRoleCmdRoleId string

var UsersAttachRoleCmdUserName string





var UsersAttachRoleCmdBody string


func init() {
  UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdOperatorId, "operator-id", "", "operator_id")

  UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdRoleId, "role-id", "", "")

  UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdUserName, "user-name", "", "user_name")



  UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  UsersCmd.AddCommand(UsersAttachRoleCmd)
}

var UsersAttachRoleCmd = &cobra.Command{
  Use: "attach-role",
  Short: TR("Attach Role to User"),
  Long: TR(`ユーザーにロールをアタッチする。`),
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
    
    param, err := collectUsersAttachRoleCmdParams()
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

func collectUsersAttachRoleCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForUsersAttachRoleCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForUsersAttachRoleCmd("/operators/{operator_id}/users/{user_name}/roles"),
    query: buildQueryForUsersAttachRoleCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForUsersAttachRoleCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersAttachRoleCmdOperatorId, -1)
  
  
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersAttachRoleCmdUserName, -1)
  
  
  
  
  
  return path
}

func buildQueryForUsersAttachRoleCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForUsersAttachRoleCmd() (string, error) {
  if UsersAttachRoleCmdBody != "" {
    if strings.HasPrefix(UsersAttachRoleCmdBody, "@") {
      fname := strings.TrimPrefix(UsersAttachRoleCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if UsersAttachRoleCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return UsersAttachRoleCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  
  
  if UsersAttachRoleCmdRoleId != "" {
    result["roleId"] = UsersAttachRoleCmdRoleId
  }
  
  
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

