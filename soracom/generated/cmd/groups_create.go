package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)










var GroupsCreateCmdBody string


func init() {


  GroupsCreateCmd.Flags().StringVar(&GroupsCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  GroupsCmd.AddCommand(GroupsCreateCmd)
}

var GroupsCreateCmd = &cobra.Command{
  Use: "create",
  Short: TR("Create Group"),
  Long: TR(`Groupを新規作成する
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
    
    param, err := collectGroupsCreateCmdParams()
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

func collectGroupsCreateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForGroupsCreateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForGroupsCreateCmd("/groups"),
    query: buildQueryForGroupsCreateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForGroupsCreateCmd(path string) string {
  
  
  
  
  return path
}

func buildQueryForGroupsCreateCmd() string {
  result := []string{}
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForGroupsCreateCmd() (string, error) {
  if GroupsCreateCmdBody != "" {
    if strings.HasPrefix(GroupsCreateCmdBody, "@") {
      fname := strings.TrimPrefix(GroupsCreateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if GroupsCreateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return GroupsCreateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

