package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail string

var SubscribersIssueTransferTokenCmdTransferDestinationOperatorId string





var SubscribersIssueTransferTokenCmdBody string


func init() {
  SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail, "transfer-destination-operator-email", "", "")

  SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdTransferDestinationOperatorId, "transfer-destination-operator-id", "", "")



  SubscribersIssueTransferTokenCmd.Flags().StringVar(&SubscribersIssueTransferTokenCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  SubscribersCmd.AddCommand(SubscribersIssueTransferTokenCmd)
}

var SubscribersIssueTransferTokenCmd = &cobra.Command{
  Use: "issue-transfer-token",
  Short: TR("Issue Subscribers Transfer Token"),
  Long: TR(`Subscriberのオペレーター間移管用トークンを、移管先オペレーターにメールで送付する。`),
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
    
    param, err := collectSubscribersIssueTransferTokenCmdParams()
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

func collectSubscribersIssueTransferTokenCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForSubscribersIssueTransferTokenCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersIssueTransferTokenCmd("/subscribers/transfer_token/issue"),
    query: buildQueryForSubscribersIssueTransferTokenCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForSubscribersIssueTransferTokenCmd(path string) string {
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForSubscribersIssueTransferTokenCmd() string {
  result := []string{}
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForSubscribersIssueTransferTokenCmd() (string, error) {
  if SubscribersIssueTransferTokenCmdBody != "" {
    if strings.HasPrefix(SubscribersIssueTransferTokenCmdBody, "@") {
      fname := strings.TrimPrefix(SubscribersIssueTransferTokenCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if SubscribersIssueTransferTokenCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return SubscribersIssueTransferTokenCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail != "" {
    result["transferDestinationOperatorEmail"] = SubscribersIssueTransferTokenCmdTransferDestinationOperatorEmail
  }
  
  
  
  if SubscribersIssueTransferTokenCmdTransferDestinationOperatorId != "" {
    result["transferDestinationOperatorId"] = SubscribersIssueTransferTokenCmdTransferDestinationOperatorId
  }
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

