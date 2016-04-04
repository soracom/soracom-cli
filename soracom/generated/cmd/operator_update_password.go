package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorUpdatePasswordCmdCurrentPassword string

var OperatorUpdatePasswordCmdNewPassword string

var OperatorUpdatePasswordCmdOperatorId string




var OperatorUpdatePasswordCmdBody string


func init() {
  OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdCurrentPassword, "current-password", "", "")

  OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdNewPassword, "new-password", "", "")

  OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdOperatorId, "operator-id", "", "operator ID")



  OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  OperatorCmd.AddCommand(OperatorUpdatePasswordCmd)
}

var OperatorUpdatePasswordCmd = &cobra.Command{
  Use: "update-password",
  Short: TR("Update Operator Password"),
  Long: TR(`Operator のパスワードを更新する。
`),
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
    
    param, err := collectOperatorUpdatePasswordCmdParams()
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

func collectOperatorUpdatePasswordCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForOperatorUpdatePasswordCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForOperatorUpdatePasswordCmd("/operators/{operator_id}/password"),
    query: buildQueryForOperatorUpdatePasswordCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForOperatorUpdatePasswordCmd(path string) string {
  
  
  
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", OperatorUpdatePasswordCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForOperatorUpdatePasswordCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForOperatorUpdatePasswordCmd() (string, error) {
  if OperatorUpdatePasswordCmdBody != "" {
    if strings.HasPrefix(OperatorUpdatePasswordCmdBody, "@") {
      fname := strings.TrimPrefix(OperatorUpdatePasswordCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if OperatorUpdatePasswordCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return OperatorUpdatePasswordCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if OperatorUpdatePasswordCmdCurrentPassword != "" {
    result["currentPassword"] = OperatorUpdatePasswordCmdCurrentPassword
  }
  
  if OperatorUpdatePasswordCmdNewPassword != "" {
    result["newPassword"] = OperatorUpdatePasswordCmdNewPassword
  }
  
  if OperatorUpdatePasswordCmdOperatorId != "" {
    result["operator_id"] = OperatorUpdatePasswordCmdOperatorId
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

