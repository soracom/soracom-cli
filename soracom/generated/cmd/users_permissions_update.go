package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersPermissionsUpdateCmdDescription string

var UsersPermissionsUpdateCmdOperatorId string

var UsersPermissionsUpdateCmdPermission string

var UsersPermissionsUpdateCmdUserName string




var UsersPermissionsUpdateCmdBody string


func init() {
  UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdDescription, "description", "", "")

  UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdOperatorId, "operator-id", "", "operator_id")

  UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdPermission, "permission", "", "")

  UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdUserName, "user-name", "", "user_name")



  UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  UsersPermissionsCmd.AddCommand(UsersPermissionsUpdateCmd)
}

var UsersPermissionsUpdateCmd = &cobra.Command{
  Use: "update",
  Short: TR("Update Permission to User"),
  Long: TR(`SAMユーザーの権限を更新する。`),
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
    
    param, err := collectUsersPermissionsUpdateCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    fmt.Println(result)
    return nil
  },
}

func collectUsersPermissionsUpdateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForUsersPermissionsUpdateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "PUT",
    path: buildPathForUsersPermissionsUpdateCmd("/operators/{operator_id}/users/{user_name}/permission"),
    query: buildQueryForUsersPermissionsUpdateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForUsersPermissionsUpdateCmd(path string) string {
  
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersPermissionsUpdateCmdOperatorId, -1)
  
  
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersPermissionsUpdateCmdUserName, -1)
  
  
  
  
  return path
}

func buildQueryForUsersPermissionsUpdateCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForUsersPermissionsUpdateCmd() (string, error) {
  if UsersPermissionsUpdateCmdBody != "" {
    if strings.HasPrefix(UsersPermissionsUpdateCmdBody, "@") {
      fname := strings.TrimPrefix(UsersPermissionsUpdateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if UsersPermissionsUpdateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return UsersPermissionsUpdateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if UsersPermissionsUpdateCmdDescription != "" {
    result["description"] = UsersPermissionsUpdateCmdDescription
  }
  
  if UsersPermissionsUpdateCmdOperatorId != "" {
    result["operator_id"] = UsersPermissionsUpdateCmdOperatorId
  }
  
  if UsersPermissionsUpdateCmdPermission != "" {
    result["permission"] = UsersPermissionsUpdateCmdPermission
  }
  
  if UsersPermissionsUpdateCmdUserName != "" {
    result["user_name"] = UsersPermissionsUpdateCmdUserName
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

