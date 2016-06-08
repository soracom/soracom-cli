package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsPutConfigCmdGroupId string

var GroupsPutConfigCmdNamespace string





var GroupsPutConfigCmdBody string


func init() {
  GroupsPutConfigCmd.Flags().StringVar(&GroupsPutConfigCmdGroupId, "group-id", "", TR("groups.put_configuration_parameters.put.parameters.group_id.description"))

  GroupsPutConfigCmd.Flags().StringVar(&GroupsPutConfigCmdNamespace, "namespace", "", TR("groups.put_configuration_parameters.put.parameters.namespace.description"))



  GroupsPutConfigCmd.Flags().StringVar(&GroupsPutConfigCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  GroupsCmd.AddCommand(GroupsPutConfigCmd)
}

var GroupsPutConfigCmd = &cobra.Command{
  Use: "put-config",
  Short: TR("groups.put_configuration_parameters.put.summary"),
  Long: TR(`groups.put_configuration_parameters.put.description`),
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
    
    param, err := collectGroupsPutConfigCmdParams()
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

func collectGroupsPutConfigCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForGroupsPutConfigCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "PUT",
    path: buildPathForGroupsPutConfigCmd("/groups/{group_id}/configuration/{namespace}"),
    query: buildQueryForGroupsPutConfigCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForGroupsPutConfigCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsPutConfigCmdGroupId, -1)
  
  
  
  path = strings.Replace(path, "{" + "namespace" + "}", GroupsPutConfigCmdNamespace, -1)
  
  
  
  
  
  return path
}

func buildQueryForGroupsPutConfigCmd() string {
  result := []string{}
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForGroupsPutConfigCmd() (string, error) {
  if GroupsPutConfigCmdBody != "" {
    if strings.HasPrefix(GroupsPutConfigCmdBody, "@") {
      fname := strings.TrimPrefix(GroupsPutConfigCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if GroupsPutConfigCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return GroupsPutConfigCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

