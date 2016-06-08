package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersUnignoreCmdHandlerId string

var EventHandlersUnignoreCmdImsi string






func init() {
  EventHandlersUnignoreCmd.Flags().StringVar(&EventHandlersUnignoreCmdHandlerId, "handler-id", "", TR("handler_id"))

  EventHandlersUnignoreCmd.Flags().StringVar(&EventHandlersUnignoreCmdImsi, "imsi", "", TR("imsi"))




  EventHandlersCmd.AddCommand(EventHandlersUnignoreCmd)
}

var EventHandlersUnignoreCmd = &cobra.Command{
  Use: "unignore",
  Short: TR("event_handlers.delete_ignore_event_handler.delete.summary"),
  Long: TR(`event_handlers.delete_ignore_event_handler.delete.description`),
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
    
    param, err := collectEventHandlersUnignoreCmdParams()
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

func collectEventHandlersUnignoreCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForEventHandlersUnignoreCmd("/event_handlers/{handler_id}/subscribers/{imsi}/ignore"),
    query: buildQueryForEventHandlersUnignoreCmd(),
    
    
  }, nil
}

func buildPathForEventHandlersUnignoreCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "handler_id" + "}", EventHandlersUnignoreCmdHandlerId, -1)
  
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", EventHandlersUnignoreCmdImsi, -1)
  
  
  
  
  
  return path
}

func buildQueryForEventHandlersUnignoreCmd() string {
  result := []string{}
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


