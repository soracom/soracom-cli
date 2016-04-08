package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorVerifyEmailChangeTokenCmdToken string




var OperatorVerifyEmailChangeTokenCmdBody string


func init() {
  OperatorVerifyEmailChangeTokenCmd.Flags().StringVar(&OperatorVerifyEmailChangeTokenCmdToken, "token", "", "")



  OperatorVerifyEmailChangeTokenCmd.Flags().StringVar(&OperatorVerifyEmailChangeTokenCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  OperatorCmd.AddCommand(OperatorVerifyEmailChangeTokenCmd)
}

var OperatorVerifyEmailChangeTokenCmd = &cobra.Command{
  Use: "verify-email-change-token",
  Short: TR("Verify Email Change Token"),
  Long: TR(`メールアドレス変更用のトークンを確認して、メールアドレスを変更する。`),
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
    
    param, err := collectOperatorVerifyEmailChangeTokenCmdParams()
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

func collectOperatorVerifyEmailChangeTokenCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForOperatorVerifyEmailChangeTokenCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForOperatorVerifyEmailChangeTokenCmd("/operators/email_change_token/verify"),
    query: buildQueryForOperatorVerifyEmailChangeTokenCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForOperatorVerifyEmailChangeTokenCmd(path string) string {
  
  
  
  
  
  return path
}

func buildQueryForOperatorVerifyEmailChangeTokenCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForOperatorVerifyEmailChangeTokenCmd() (string, error) {
  if OperatorVerifyEmailChangeTokenCmdBody != "" {
    if strings.HasPrefix(OperatorVerifyEmailChangeTokenCmdBody, "@") {
      fname := strings.TrimPrefix(OperatorVerifyEmailChangeTokenCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if OperatorVerifyEmailChangeTokenCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return OperatorVerifyEmailChangeTokenCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if OperatorVerifyEmailChangeTokenCmdToken != "" {
    result["token"] = OperatorVerifyEmailChangeTokenCmdToken
  }
  
  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

