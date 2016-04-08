package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var AuthCmdAuthKey string

var AuthCmdAuthKeyId string

var AuthCmdEmail string

var AuthCmdOperatorId string

var AuthCmdPassword string

var AuthCmdUserName string


var AuthCmdTokenTimeoutSeconds int64



var AuthCmdBody string


func init() {
  AuthCmd.Flags().StringVar(&AuthCmdAuthKey, "auth-key", "", "")

  AuthCmd.Flags().StringVar(&AuthCmdAuthKeyId, "auth-key-id", "", "")

  AuthCmd.Flags().StringVar(&AuthCmdEmail, "email", "", "")

  AuthCmd.Flags().StringVar(&AuthCmdOperatorId, "operator-id", "", "")

  AuthCmd.Flags().StringVar(&AuthCmdPassword, "password", "", "")

  AuthCmd.Flags().StringVar(&AuthCmdUserName, "user-name", "", "")

  AuthCmd.Flags().Int64Var(&AuthCmdTokenTimeoutSeconds, "token-timeout-seconds", 0, "")



  AuthCmd.Flags().StringVar(&AuthCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  RootCmd.AddCommand(AuthCmd)
}

var AuthCmd = &cobra.Command{
  Use: "auth",
  Short: TR("auth.post.summary"),
  Long: TR(`auth.post.description`),
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

    
    param, err := collectAuthCmdParams()
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

func collectAuthCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForAuthCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForAuthCmd("/auth"),
    query: buildQueryForAuthCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForAuthCmd(path string) string {
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForAuthCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  
  
  
  
  

  
  
  

  
  return strings.Join(result, "&")
}


func buildBodyForAuthCmd() (string, error) {
  if AuthCmdBody != "" {
    if strings.HasPrefix(AuthCmdBody, "@") {
      fname := strings.TrimPrefix(AuthCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if AuthCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return AuthCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if AuthCmdAuthKey != "" {
    result["authKey"] = AuthCmdAuthKey
  }
  
  
  
  if AuthCmdAuthKeyId != "" {
    result["authKeyId"] = AuthCmdAuthKeyId
  }
  
  
  
  if AuthCmdEmail != "" {
    result["email"] = AuthCmdEmail
  }
  
  
  
  if AuthCmdOperatorId != "" {
    result["operatorId"] = AuthCmdOperatorId
  }
  
  
  
  if AuthCmdPassword != "" {
    result["password"] = AuthCmdPassword
  }
  
  
  
  if AuthCmdUserName != "" {
    result["userName"] = AuthCmdUserName
  }
  
  

  
  
  if AuthCmdTokenTimeoutSeconds != 0 {
    result["tokenTimeoutSeconds"] = AuthCmdTokenTimeoutSeconds
  }
  
  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

