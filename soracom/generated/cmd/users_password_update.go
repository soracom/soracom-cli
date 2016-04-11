package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersPasswordUpdateCmdCurrentPassword string

var UsersPasswordUpdateCmdNewPassword string

var UsersPasswordUpdateCmdOperatorId string

var UsersPasswordUpdateCmdUserName string





var UsersPasswordUpdateCmdBody string


func init() {
  UsersPasswordUpdateCmd.Flags().StringVar(&UsersPasswordUpdateCmdCurrentPassword, "current-password", "", "")

  UsersPasswordUpdateCmd.Flags().StringVar(&UsersPasswordUpdateCmdNewPassword, "new-password", "", "")

  UsersPasswordUpdateCmd.Flags().StringVar(&UsersPasswordUpdateCmdOperatorId, "operator-id", "", "operator_id")

  UsersPasswordUpdateCmd.Flags().StringVar(&UsersPasswordUpdateCmdUserName, "user-name", "", "user_name")



  UsersPasswordUpdateCmd.Flags().StringVar(&UsersPasswordUpdateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  UsersPasswordCmd.AddCommand(UsersPasswordUpdateCmd)
}

var UsersPasswordUpdateCmd = &cobra.Command{
  Use: "update",
  Short: TR("Update Password"),
  Long: TR(`SAMユーザーのパスワードを更新する。`),
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
    
    param, err := collectUsersPasswordUpdateCmdParams()
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

func collectUsersPasswordUpdateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForUsersPasswordUpdateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "PUT",
    path: buildPathForUsersPasswordUpdateCmd("/operators/{operator_id}/users/{user_name}/password"),
    query: buildQueryForUsersPasswordUpdateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForUsersPasswordUpdateCmd(path string) string {
  
  
  
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersPasswordUpdateCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersPasswordUpdateCmdUserName, -1)
  
  
  
  
  
  return path
}

func buildQueryForUsersPasswordUpdateCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForUsersPasswordUpdateCmd() (string, error) {
  if UsersPasswordUpdateCmdBody != "" {
    if strings.HasPrefix(UsersPasswordUpdateCmdBody, "@") {
      fname := strings.TrimPrefix(UsersPasswordUpdateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if UsersPasswordUpdateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return UsersPasswordUpdateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if UsersPasswordUpdateCmdCurrentPassword != "" {
    result["currentPassword"] = UsersPasswordUpdateCmdCurrentPassword
  }
  
  
  
  if UsersPasswordUpdateCmdNewPassword != "" {
    result["newPassword"] = UsersPasswordUpdateCmdNewPassword
  }
  
  
  
  
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

