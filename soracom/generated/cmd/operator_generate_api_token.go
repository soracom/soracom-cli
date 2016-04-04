package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorGenerateApiTokenCmdOperatorId string




var OperatorGenerateApiTokenCmdBody string


func init() {
  OperatorGenerateApiTokenCmd.Flags().StringVar(&OperatorGenerateApiTokenCmdOperatorId, "operator-id", "", "operator ID")



  OperatorGenerateApiTokenCmd.Flags().StringVar(&OperatorGenerateApiTokenCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  OperatorCmd.AddCommand(OperatorGenerateApiTokenCmd)
}

var OperatorGenerateApiTokenCmd = &cobra.Command{
  Use: "generate-api-token",
  Short: TR("Generate Authentication Token"),
  Long: TR(`新しい API トークンを発行する。
現在の API トークンをヘッダーに入れてリクエストを行うと、新しい API トークンを含んだレスポンスを返す。
以降のリクエスト時にはこの新しい API トークンを用いることができる。
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
    
    param, err := collectOperatorGenerateApiTokenCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    fmt.Println(result)
    return nil
  },
}

func collectOperatorGenerateApiTokenCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForOperatorGenerateApiTokenCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForOperatorGenerateApiTokenCmd("/operators/{operator_id}/token"),
    query: buildQueryForOperatorGenerateApiTokenCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForOperatorGenerateApiTokenCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", OperatorGenerateApiTokenCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForOperatorGenerateApiTokenCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForOperatorGenerateApiTokenCmd() (string, error) {
  if OperatorGenerateApiTokenCmdBody != "" {
    if strings.HasPrefix(OperatorGenerateApiTokenCmdBody, "@") {
      fname := strings.TrimPrefix(OperatorGenerateApiTokenCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if OperatorGenerateApiTokenCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return OperatorGenerateApiTokenCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if OperatorGenerateApiTokenCmdOperatorId != "" {
    result["operator_id"] = OperatorGenerateApiTokenCmdOperatorId
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

