package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var AuthIssuePasswordResetTokenCmdEmail string




var AuthIssuePasswordResetTokenCmdBody string


func init() {
  AuthIssuePasswordResetTokenCmd.Flags().StringVar(&AuthIssuePasswordResetTokenCmdEmail, "email", "", "")



  AuthIssuePasswordResetTokenCmd.Flags().StringVar(&AuthIssuePasswordResetTokenCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  AuthCmd.AddCommand(AuthIssuePasswordResetTokenCmd)
}

var AuthIssuePasswordResetTokenCmd = &cobra.Command{
  Use: "issue-password-reset-token",
  Short: TR("auth.issue_password_reset_token.post.summary"),
  Long: TR(`auth.issue_password_reset_token.post.description`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
    }

    ac := newAPIClient(opt)
    if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
      ac.SetVerbose(true)
    }

    
    param, err := collectAuthIssuePasswordResetTokenCmdParams()
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

func collectAuthIssuePasswordResetTokenCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForAuthIssuePasswordResetTokenCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForAuthIssuePasswordResetTokenCmd("/auth/password_reset_token/issue"),
    query: buildQueryForAuthIssuePasswordResetTokenCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForAuthIssuePasswordResetTokenCmd(path string) string {
  
  
  
  
  
  return path
}

func buildQueryForAuthIssuePasswordResetTokenCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForAuthIssuePasswordResetTokenCmd() (string, error) {
  if AuthIssuePasswordResetTokenCmdBody != "" {
    if strings.HasPrefix(AuthIssuePasswordResetTokenCmdBody, "@") {
      fname := strings.TrimPrefix(AuthIssuePasswordResetTokenCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if AuthIssuePasswordResetTokenCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return AuthIssuePasswordResetTokenCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if AuthIssuePasswordResetTokenCmdEmail != "" {
    result["email"] = AuthIssuePasswordResetTokenCmdEmail
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

