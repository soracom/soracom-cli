package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var RolesUpdateCmdDescription string

var RolesUpdateCmdOperatorId string

var RolesUpdateCmdPermission string

var RolesUpdateCmdRoleId string





var RolesUpdateCmdBody string


func init() {
  RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdDescription, "description", "", "")

  RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdOperatorId, "operator-id", "", "operator_id")

  RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdPermission, "permission", "", "")

  RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdRoleId, "role-id", "", "role_id")



  RolesUpdateCmd.Flags().StringVar(&RolesUpdateCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  RolesCmd.AddCommand(RolesUpdateCmd)
}

var RolesUpdateCmd = &cobra.Command{
  Use: "update",
  Short: TR("Update Role"),
  Long: TR(`Roleを編集する。`),
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
    
    param, err := collectRolesUpdateCmdParams()
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

func collectRolesUpdateCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForRolesUpdateCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "PUT",
    path: buildPathForRolesUpdateCmd("/operators/{operator_id}/roles/{role_id}"),
    query: buildQueryForRolesUpdateCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForRolesUpdateCmd(path string) string {
  
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", RolesUpdateCmdOperatorId, -1)
  
  
  
  
  
  path = strings.Replace(path, "{" + "role_id" + "}", RolesUpdateCmdRoleId, -1)
  
  
  
  
  
  return path
}

func buildQueryForRolesUpdateCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForRolesUpdateCmd() (string, error) {
  if RolesUpdateCmdBody != "" {
    if strings.HasPrefix(RolesUpdateCmdBody, "@") {
      fname := strings.TrimPrefix(RolesUpdateCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if RolesUpdateCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return RolesUpdateCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if RolesUpdateCmdDescription != "" {
    result["description"] = RolesUpdateCmdDescription
  }
  
  
  
  
  
  if RolesUpdateCmdPermission != "" {
    result["permission"] = RolesUpdateCmdPermission
  }
  
  
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

