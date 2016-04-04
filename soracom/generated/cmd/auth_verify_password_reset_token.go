package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var AuthVerifyPasswordResetTokenCmdPassword string

var AuthVerifyPasswordResetTokenCmdToken string




var AuthVerifyPasswordResetTokenCmdBody string


func init() {
  AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdPassword, "password", "", "")

  AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdToken, "token", "", "")



  AuthVerifyPasswordResetTokenCmd.Flags().StringVar(&AuthVerifyPasswordResetTokenCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  AuthCmd.AddCommand(AuthVerifyPasswordResetTokenCmd)
}

var AuthVerifyPasswordResetTokenCmd = &cobra.Command{
  Use: "verify-password-reset-token",
  Short: TR("auth.verify_password_reset_token.post.summary"),
  Long: TR(`auth.verify_password_reset_token.post.description`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
    }

    ac := newAPIClient(opt)
    if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
      ac.SetVerbose(true)
    }

    
    param, err := collectAuthVerifyPasswordResetTokenCmdParams()
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

func collectAuthVerifyPasswordResetTokenCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForAuthVerifyPasswordResetTokenCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForAuthVerifyPasswordResetTokenCmd("/auth/password_reset_token/verify"),
    query: buildQueryForAuthVerifyPasswordResetTokenCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForAuthVerifyPasswordResetTokenCmd(path string) string {
  
  
  
  
  
  
  
  return path
}

func buildQueryForAuthVerifyPasswordResetTokenCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForAuthVerifyPasswordResetTokenCmd() (string, error) {
  if AuthVerifyPasswordResetTokenCmdBody != "" {
    if strings.HasPrefix(AuthVerifyPasswordResetTokenCmdBody, "@") {
      fname := strings.TrimPrefix(AuthVerifyPasswordResetTokenCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if AuthVerifyPasswordResetTokenCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return AuthVerifyPasswordResetTokenCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if AuthVerifyPasswordResetTokenCmdPassword != "" {
    result["password"] = AuthVerifyPasswordResetTokenCmdPassword
  }
  
  if AuthVerifyPasswordResetTokenCmdToken != "" {
    result["token"] = AuthVerifyPasswordResetTokenCmdToken
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

