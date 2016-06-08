package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersDeleteTransferTokenCmdToken string






func init() {
  SubscribersDeleteTransferTokenCmd.Flags().StringVar(&SubscribersDeleteTransferTokenCmdToken, "token", "", TR("subscribers.delete_subscriber_transfer_token.delete.parameters.token.description"))




  SubscribersCmd.AddCommand(SubscribersDeleteTransferTokenCmd)
}

var SubscribersDeleteTransferTokenCmd = &cobra.Command{
  Use: "delete-transfer-token",
  Short: TR("subscribers.delete_subscriber_transfer_token.delete.summary"),
  Long: TR(`subscribers.delete_subscriber_transfer_token.delete.description`),
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
    
    param, err := collectSubscribersDeleteTransferTokenCmdParams()
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

func collectSubscribersDeleteTransferTokenCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForSubscribersDeleteTransferTokenCmd("/subscribers/transfer_token/{token}"),
    query: buildQueryForSubscribersDeleteTransferTokenCmd(),
    
    
  }, nil
}

func buildPathForSubscribersDeleteTransferTokenCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "token" + "}", SubscribersDeleteTransferTokenCmdToken, -1)
  
  
  
  
  
  return path
}

func buildQueryForSubscribersDeleteTransferTokenCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


