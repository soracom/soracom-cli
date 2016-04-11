package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersSetGroupCmdGroupId string

var SubscribersSetGroupCmdImsi string

var SubscribersSetGroupCmdOperatorId string


var SubscribersSetGroupCmdCreatedTime int64

var SubscribersSetGroupCmdLastModifiedTime int64




var SubscribersSetGroupCmdBody string


func init() {
  SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdGroupId, "group-id", "", "")

  SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdImsi, "imsi", "", "対象のSubscriberのIMSI")

  SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdOperatorId, "operator-id", "", "")

  SubscribersSetGroupCmd.Flags().Int64Var(&SubscribersSetGroupCmdCreatedTime, "created-time", 0, "")

  SubscribersSetGroupCmd.Flags().Int64Var(&SubscribersSetGroupCmdLastModifiedTime, "last-modified-time", 0, "")



  SubscribersSetGroupCmd.Flags().StringVar(&SubscribersSetGroupCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  SubscribersCmd.AddCommand(SubscribersSetGroupCmd)
}

var SubscribersSetGroupCmd = &cobra.Command{
  Use: "set-group",
  Short: TR("Set Group to Subscriber"),
  Long: TR(`指定されたSubscriberの所属先Groupを指定あるいは上書き変更`),
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
    
    param, err := collectSubscribersSetGroupCmdParams()
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

func collectSubscribersSetGroupCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForSubscribersSetGroupCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersSetGroupCmd("/subscribers/{imsi}/set_group"),
    query: buildQueryForSubscribersSetGroupCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForSubscribersSetGroupCmd(path string) string {
  
  
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersSetGroupCmdImsi, -1)
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForSubscribersSetGroupCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  
  
  
  
  

  

  

  return strings.Join(result, "&")
}


func buildBodyForSubscribersSetGroupCmd() (string, error) {
  if SubscribersSetGroupCmdBody != "" {
    if strings.HasPrefix(SubscribersSetGroupCmdBody, "@") {
      fname := strings.TrimPrefix(SubscribersSetGroupCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if SubscribersSetGroupCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return SubscribersSetGroupCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if SubscribersSetGroupCmdGroupId != "" {
    result["groupId"] = SubscribersSetGroupCmdGroupId
  }
  
  
  
  
  
  if SubscribersSetGroupCmdOperatorId != "" {
    result["operatorId"] = SubscribersSetGroupCmdOperatorId
  }
  
  

  
  
  if SubscribersSetGroupCmdCreatedTime != 0 {
    result["createdTime"] = SubscribersSetGroupCmdCreatedTime
  }
  
  
  
  if SubscribersSetGroupCmdLastModifiedTime != 0 {
    result["lastModifiedTime"] = SubscribersSetGroupCmdLastModifiedTime
  }
  
  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

