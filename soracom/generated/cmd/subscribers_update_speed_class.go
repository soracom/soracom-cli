package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersUpdateSpeedClassCmdImsi string

var SubscribersUpdateSpeedClassCmdSpeedClass string




var SubscribersUpdateSpeedClassCmdBody string


func init() {
  SubscribersUpdateSpeedClassCmd.Flags().StringVar(&SubscribersUpdateSpeedClassCmdImsi, "imsi", "", "対象のSubscriberのIMSI")

  SubscribersUpdateSpeedClassCmd.Flags().StringVar(&SubscribersUpdateSpeedClassCmdSpeedClass, "speed-class", "", "")



  SubscribersUpdateSpeedClassCmd.Flags().StringVar(&SubscribersUpdateSpeedClassCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  SubscribersCmd.AddCommand(SubscribersUpdateSpeedClassCmd)
}

var SubscribersUpdateSpeedClassCmd = &cobra.Command{
  Use: "update-speed-class",
  Short: TR("Update Subscriber speed class"),
  Long: TR(`指定されたSubscriberの速度クラスを変更します`),
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
    
    param, err := collectSubscribersUpdateSpeedClassCmdParams()
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

func collectSubscribersUpdateSpeedClassCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForSubscribersUpdateSpeedClassCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersUpdateSpeedClassCmd("/subscribers/{imsi}/update_speed_class"),
    query: buildQueryForSubscribersUpdateSpeedClassCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForSubscribersUpdateSpeedClassCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersUpdateSpeedClassCmdImsi, -1)
  
  
  
  
  
  
  return path
}

func buildQueryForSubscribersUpdateSpeedClassCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForSubscribersUpdateSpeedClassCmd() (string, error) {
  if SubscribersUpdateSpeedClassCmdBody != "" {
    if strings.HasPrefix(SubscribersUpdateSpeedClassCmdBody, "@") {
      fname := strings.TrimPrefix(SubscribersUpdateSpeedClassCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if SubscribersUpdateSpeedClassCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return SubscribersUpdateSpeedClassCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if SubscribersUpdateSpeedClassCmdImsi != "" {
    result["imsi"] = SubscribersUpdateSpeedClassCmdImsi
  }
  
  if SubscribersUpdateSpeedClassCmdSpeedClass != "" {
    result["speedClass"] = SubscribersUpdateSpeedClassCmdSpeedClass
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

