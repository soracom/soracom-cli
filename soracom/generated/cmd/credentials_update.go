package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var CredentialsUpdateCmdCredentialsId string

var CredentialsUpdateCmdDescription string

var CredentialsUpdateCmdType string





var CredentialsUpdateCmdBody string


func init() {
  CredentialsUpdateCmd.Flags().StringVar(&CredentialsUpdateCmdCredentialsId, "credentials-id", "", "credentials_id")

  CredentialsUpdateCmd.Flags().StringVar(&CredentialsUpdateCmdDescription, "description", "", "")

  CredentialsUpdateCmd.Flags().StringVar(&CredentialsUpdateCmdType, "type", "", "")



  CredentialsUpdateCmd.Flags().StringVar(&CredentialsUpdateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  CredentialsCmd.AddCommand(CredentialsUpdateCmd)
}

var CredentialsUpdateCmd = &cobra.Command{
  Use: "update",
  Short: TR("Update Credential"),
  Long: TR(`認証情報を更新する。`),
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
    
    param, err := collectCredentialsUpdateCmdParams()
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

func collectCredentialsUpdateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForCredentialsUpdateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "PUT",
    path: buildPathForCredentialsUpdateCmd("/credentials/{credentials_id}"),
    query: buildQueryForCredentialsUpdateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForCredentialsUpdateCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "credentials_id" + "}", CredentialsUpdateCmdCredentialsId, -1)
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForCredentialsUpdateCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForCredentialsUpdateCmd() (string, error) {
  if CredentialsUpdateCmdBody != "" {
    if strings.HasPrefix(CredentialsUpdateCmdBody, "@") {
      fname := strings.TrimPrefix(CredentialsUpdateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if CredentialsUpdateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return CredentialsUpdateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  
  
  if CredentialsUpdateCmdDescription != "" {
    result["description"] = CredentialsUpdateCmdDescription
  }
  
  
  
  if CredentialsUpdateCmdType != "" {
    result["type"] = CredentialsUpdateCmdType
  }
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

