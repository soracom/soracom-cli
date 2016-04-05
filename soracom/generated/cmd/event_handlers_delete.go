package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersDeleteCmdHandlerId string





func init() {
  EventHandlersDeleteCmd.Flags().StringVar(&EventHandlersDeleteCmdHandlerId, "handler-id", "", "handler ID")




  EventHandlersCmd.AddCommand(EventHandlersDeleteCmd)
}

var EventHandlersDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("Delete Event Handler"),
  Long: TR(`指定されたイベントハンドラを削除する。
`),
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
    
    param, err := collectEventHandlersDeleteCmdParams()
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

func collectEventHandlersDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForEventHandlersDeleteCmd("/event_handlers/{handler_id}"),
    query: buildQueryForEventHandlersDeleteCmd(),
    
    
  }, nil
}

func buildPathForEventHandlersDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "handler_id" + "}", EventHandlersDeleteCmdHandlerId, -1)
  
  
  
  
  return path
}

func buildQueryForEventHandlersDeleteCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


