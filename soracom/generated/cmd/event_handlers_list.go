package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersListCmdTarget string





func init() {
  EventHandlersListCmd.Flags().StringVar(&EventHandlersListCmdTarget, "target", "", "target")




  EventHandlersCmd.AddCommand(EventHandlersListCmd)
}

var EventHandlersListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List Event Handlers"),
  Long: TR(`イベントハンドラの一覧を返す。
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
    
    param, err := collectEventHandlersListCmdParams()
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

func collectEventHandlersListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForEventHandlersListCmd("/event_handlers"),
    query: buildQueryForEventHandlersListCmd(),
    
    
  }, nil
}

func buildPathForEventHandlersListCmd(path string) string {
  
  
  
  
  
  return path
}

func buildQueryForEventHandlersListCmd() string {
  result := []string{}
  
  
  if EventHandlersListCmdTarget != "" {
    result = append(result, sprintf("%s=%s", "target", EventHandlersListCmdTarget))
  }
  
  

  

  
  return strings.Join(result, "&")
}


