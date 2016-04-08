package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersRegisterCmdGroupId string

var SubscribersRegisterCmdImsi string

var SubscribersRegisterCmdRegistrationSecret string




var SubscribersRegisterCmdBody string


func init() {
  SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdGroupId, "group-id", "", "")

  SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdImsi, "imsi", "", "対象のSubscriberのIMSI")

  SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdRegistrationSecret, "registration-secret", "", "")



  SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  SubscribersCmd.AddCommand(SubscribersRegisterCmd)
}

var SubscribersRegisterCmd = &cobra.Command{
  Use: "register",
  Short: TR("Register Subscriber"),
  Long: TR(`Subscriberを登録`),
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
    
    param, err := collectSubscribersRegisterCmdParams()
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

func collectSubscribersRegisterCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForSubscribersRegisterCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersRegisterCmd("/subscribers/{imsi}/register"),
    query: buildQueryForSubscribersRegisterCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForSubscribersRegisterCmd(path string) string {
  
  
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersRegisterCmdImsi, -1)
  
  
  
  
  
  
  return path
}

func buildQueryForSubscribersRegisterCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForSubscribersRegisterCmd() (string, error) {
  if SubscribersRegisterCmdBody != "" {
    if strings.HasPrefix(SubscribersRegisterCmdBody, "@") {
      fname := strings.TrimPrefix(SubscribersRegisterCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if SubscribersRegisterCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return SubscribersRegisterCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if SubscribersRegisterCmdGroupId != "" {
    result["groupId"] = SubscribersRegisterCmdGroupId
  }
  
  
  
  
  
  if SubscribersRegisterCmdRegistrationSecret != "" {
    result["registrationSecret"] = SubscribersRegisterCmdRegistrationSecret
  }
  
  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

