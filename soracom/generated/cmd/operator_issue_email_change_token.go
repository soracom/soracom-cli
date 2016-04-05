package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorIssueEmailChangeTokenCmdEmail string




var OperatorIssueEmailChangeTokenCmdBody string


func init() {
  OperatorIssueEmailChangeTokenCmd.Flags().StringVar(&OperatorIssueEmailChangeTokenCmdEmail, "email", "", "")



  OperatorIssueEmailChangeTokenCmd.Flags().StringVar(&OperatorIssueEmailChangeTokenCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  OperatorCmd.AddCommand(OperatorIssueEmailChangeTokenCmd)
}

var OperatorIssueEmailChangeTokenCmd = &cobra.Command{
  Use: "issue-email-change-token",
  Short: TR("Issue Email Change Token"),
  Long: TR(`メールアドレス変更用のワンタイムトークンをメールで送付する。`),
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
    
    param, err := collectOperatorIssueEmailChangeTokenCmdParams()
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

func collectOperatorIssueEmailChangeTokenCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForOperatorIssueEmailChangeTokenCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForOperatorIssueEmailChangeTokenCmd("/operators/email_change_token/issue"),
    query: buildQueryForOperatorIssueEmailChangeTokenCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForOperatorIssueEmailChangeTokenCmd(path string) string {
  
  
  
  
  
  return path
}

func buildQueryForOperatorIssueEmailChangeTokenCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForOperatorIssueEmailChangeTokenCmd() (string, error) {
  if OperatorIssueEmailChangeTokenCmdBody != "" {
    if strings.HasPrefix(OperatorIssueEmailChangeTokenCmdBody, "@") {
      fname := strings.TrimPrefix(OperatorIssueEmailChangeTokenCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if OperatorIssueEmailChangeTokenCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return OperatorIssueEmailChangeTokenCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if OperatorIssueEmailChangeTokenCmdEmail != "" {
    result["email"] = OperatorIssueEmailChangeTokenCmdEmail
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}
