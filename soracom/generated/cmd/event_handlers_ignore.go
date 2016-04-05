package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersIgnoreCmdHandlerId string

var EventHandlersIgnoreCmdImsi string





func init() {
  EventHandlersIgnoreCmd.Flags().StringVar(&EventHandlersIgnoreCmdHandlerId, "handler-id", "", "handler_id")

  EventHandlersIgnoreCmd.Flags().StringVar(&EventHandlersIgnoreCmdImsi, "imsi", "", "imsi")




  EventHandlersCmd.AddCommand(EventHandlersIgnoreCmd)
}

var EventHandlersIgnoreCmd = &cobra.Command{
  Use: "ignore",
  Short: TR("Ignore Event Handler"),
  Long: TR(`指定のIMSIに対して、指定のイベントハンドラを無視する設定を追加`),
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
    
    param, err := collectEventHandlersIgnoreCmdParams()
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

func collectEventHandlersIgnoreCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForEventHandlersIgnoreCmd("/event_handlers/{handler_id}/subscribers/{imsi}/ignore"),
    query: buildQueryForEventHandlersIgnoreCmd(),
    
    
  }, nil
}

func buildPathForEventHandlersIgnoreCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "handler_id" + "}", EventHandlersIgnoreCmdHandlerId, -1)
  
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", EventHandlersIgnoreCmdImsi, -1)
  
  
  
  
  return path
}

func buildQueryForEventHandlersIgnoreCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


