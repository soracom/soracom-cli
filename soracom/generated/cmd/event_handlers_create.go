package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersCreateCmdDescription string

var EventHandlersCreateCmdName string

var EventHandlersCreateCmdStatus string

var EventHandlersCreateCmdTargetGroupId string

var EventHandlersCreateCmdTargetImsi string

var EventHandlersCreateCmdTargetOperatorId string





var EventHandlersCreateCmdBody string


func init() {
  EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdDescription, "description", "", "")

  EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdName, "name", "", "")

  EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdStatus, "status", "", "")

  EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetGroupId, "target-group-id", "", "")

  EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetImsi, "target-imsi", "", "")

  EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdTargetOperatorId, "target-operator-id", "", "")



  EventHandlersCreateCmd.Flags().StringVar(&EventHandlersCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  EventHandlersCmd.AddCommand(EventHandlersCreateCmd)
}

var EventHandlersCreateCmd = &cobra.Command{
  Use: "create",
  Short: TR("Create Event Handler"),
  Long: TR(`イベントハンドラを新規作成する。
`),
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
    
    param, err := collectEventHandlersCreateCmdParams()
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

func collectEventHandlersCreateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForEventHandlersCreateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForEventHandlersCreateCmd("/event_handlers"),
    query: buildQueryForEventHandlersCreateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForEventHandlersCreateCmd(path string) string {
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForEventHandlersCreateCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForEventHandlersCreateCmd() (string, error) {
  if EventHandlersCreateCmdBody != "" {
    if strings.HasPrefix(EventHandlersCreateCmdBody, "@") {
      fname := strings.TrimPrefix(EventHandlersCreateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if EventHandlersCreateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return EventHandlersCreateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if EventHandlersCreateCmdDescription != "" {
    result["description"] = EventHandlersCreateCmdDescription
  }
  
  
  
  if EventHandlersCreateCmdName != "" {
    result["name"] = EventHandlersCreateCmdName
  }
  
  
  
  if EventHandlersCreateCmdStatus != "" {
    result["status"] = EventHandlersCreateCmdStatus
  }
  
  
  
  if EventHandlersCreateCmdTargetGroupId != "" {
    result["targetGroupId"] = EventHandlersCreateCmdTargetGroupId
  }
  
  
  
  if EventHandlersCreateCmdTargetImsi != "" {
    result["targetImsi"] = EventHandlersCreateCmdTargetImsi
  }
  
  
  
  if EventHandlersCreateCmdTargetOperatorId != "" {
    result["targetOperatorId"] = EventHandlersCreateCmdTargetOperatorId
  }
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

