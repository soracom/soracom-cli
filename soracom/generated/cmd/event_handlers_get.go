package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersGetCmdHandlerId string






func init() {
  EventHandlersGetCmd.Flags().StringVar(&EventHandlersGetCmdHandlerId, "handler-id", "", "handler ID")




  EventHandlersCmd.AddCommand(EventHandlersGetCmd)
}

var EventHandlersGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get Event Handler"),
  Long: TR(`指定されたイベントハンドラの情報を返す。
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
    
    param, err := collectEventHandlersGetCmdParams()
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

func collectEventHandlersGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForEventHandlersGetCmd("/event_handlers/{handler_id}"),
    query: buildQueryForEventHandlersGetCmd(),
    
    
  }, nil
}

func buildPathForEventHandlersGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "handler_id" + "}", EventHandlersGetCmdHandlerId, -1)
  
  
  
  
  
  return path
}

func buildQueryForEventHandlersGetCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


