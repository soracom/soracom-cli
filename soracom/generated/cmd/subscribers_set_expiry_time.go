package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersSetExpiryTimeCmdExpiryAction string

var SubscribersSetExpiryTimeCmdImsi string




var SubscribersSetExpiryTimeCmdBody string


func init() {
  SubscribersSetExpiryTimeCmd.Flags().StringVar(&SubscribersSetExpiryTimeCmdExpiryAction, "expiry-action", "", "")

  SubscribersSetExpiryTimeCmd.Flags().StringVar(&SubscribersSetExpiryTimeCmdImsi, "imsi", "", "対象のSubscriberのIMSI")



  SubscribersSetExpiryTimeCmd.Flags().StringVar(&SubscribersSetExpiryTimeCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  SubscribersCmd.AddCommand(SubscribersSetExpiryTimeCmd)
}

var SubscribersSetExpiryTimeCmd = &cobra.Command{
  Use: "set-expiry-time",
  Short: TR("Update Expiry Time of Subscriber"),
  Long: TR(`指定されたSubscriberの有効期限を更新`),
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
    
    param, err := collectSubscribersSetExpiryTimeCmdParams()
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

func collectSubscribersSetExpiryTimeCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForSubscribersSetExpiryTimeCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersSetExpiryTimeCmd("/subscribers/{imsi}/set_expiry_time"),
    query: buildQueryForSubscribersSetExpiryTimeCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForSubscribersSetExpiryTimeCmd(path string) string {
  
  
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersSetExpiryTimeCmdImsi, -1)
  
  
  
  
  return path
}

func buildQueryForSubscribersSetExpiryTimeCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForSubscribersSetExpiryTimeCmd() (string, error) {
  if SubscribersSetExpiryTimeCmdBody != "" {
    if strings.HasPrefix(SubscribersSetExpiryTimeCmdBody, "@") {
      fname := strings.TrimPrefix(SubscribersSetExpiryTimeCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if SubscribersSetExpiryTimeCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return SubscribersSetExpiryTimeCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if SubscribersSetExpiryTimeCmdExpiryAction != "" {
    result["expiryAction"] = SubscribersSetExpiryTimeCmdExpiryAction
  }
  
  if SubscribersSetExpiryTimeCmdImsi != "" {
    result["imsi"] = SubscribersSetExpiryTimeCmdImsi
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

