package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorVerifyCmdToken string





var OperatorVerifyCmdBody string


func init() {
  OperatorVerifyCmd.Flags().StringVar(&OperatorVerifyCmdToken, "token", "", TR(""))



  OperatorVerifyCmd.Flags().StringVar(&OperatorVerifyCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  OperatorCmd.AddCommand(OperatorVerifyCmd)
}

var OperatorVerifyCmd = &cobra.Command{
  Use: "verify",
  Short: TR("operator.verify_operator.post.summary"),
  Long: TR(`operator.verify_operator.post.description`),
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

    
    param, err := collectOperatorVerifyCmdParams()
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

func collectOperatorVerifyCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForOperatorVerifyCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForOperatorVerifyCmd("/operators/verify"),
    query: buildQueryForOperatorVerifyCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForOperatorVerifyCmd(path string) string {
  
  
  
  
  
  
  return path
}

func buildQueryForOperatorVerifyCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForOperatorVerifyCmd() (string, error) {
  if OperatorVerifyCmdBody != "" {
    if strings.HasPrefix(OperatorVerifyCmdBody, "@") {
      fname := strings.TrimPrefix(OperatorVerifyCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if OperatorVerifyCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return OperatorVerifyCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if OperatorVerifyCmdToken != "" {
    result["token"] = OperatorVerifyCmdToken
  }
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

