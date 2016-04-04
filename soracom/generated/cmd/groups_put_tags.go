package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsPutTagsCmdGroupId string




var GroupsPutTagsCmdBody string


func init() {
  GroupsPutTagsCmd.Flags().StringVar(&GroupsPutTagsCmdGroupId, "group-id", "", "対象のGroupのID")



  GroupsPutTagsCmd.Flags().StringVar(&GroupsPutTagsCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  GroupsCmd.AddCommand(GroupsPutTagsCmd)
}

var GroupsPutTagsCmd = &cobra.Command{
  Use: "put-tags",
  Short: TR("Update Group Tags"),
  Long: TR(`指定されたConfiguratio Groupのタグを追加・更新
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
    
    param, err := collectGroupsPutTagsCmdParams()
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

func collectGroupsPutTagsCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForGroupsPutTagsCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "PUT",
    path: buildPathForGroupsPutTagsCmd("/groups/{group_id}/tags"),
    query: buildQueryForGroupsPutTagsCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForGroupsPutTagsCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsPutTagsCmdGroupId, -1)
  
  
  
  
  return path
}

func buildQueryForGroupsPutTagsCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForGroupsPutTagsCmd() (string, error) {
  if GroupsPutTagsCmdBody != "" {
    if strings.HasPrefix(GroupsPutTagsCmdBody, "@") {
      fname := strings.TrimPrefix(GroupsPutTagsCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if GroupsPutTagsCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return GroupsPutTagsCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if GroupsPutTagsCmdGroupId != "" {
    result["group_id"] = GroupsPutTagsCmdGroupId
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

