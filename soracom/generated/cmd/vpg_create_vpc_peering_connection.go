package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock string

var VpgCreateVpcPeeringConnectionCmdPeerOwnerId string

var VpgCreateVpcPeeringConnectionCmdPeerVpcId string

var VpgCreateVpcPeeringConnectionCmdVpgId string





var VpgCreateVpcPeeringConnectionCmdBody string


func init() {
  VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock, "destination-cidr-block", "", "")

  VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdPeerOwnerId, "peer-owner-id", "", "")

  VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdPeerVpcId, "peer-vpc-id", "", "")

  VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdVpgId, "vpg-id", "", "対象のVPGのID")



  VpgCreateVpcPeeringConnectionCmd.Flags().StringVar(&VpgCreateVpcPeeringConnectionCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  VpgCmd.AddCommand(VpgCreateVpcPeeringConnectionCmd)
}

var VpgCreateVpcPeeringConnectionCmd = &cobra.Command{
  Use: "create-vpc-peering-connection",
  Short: TR("Create VPC Peering Connection"),
  Long: TR(`指定されたVPGにVPC Peering Connectionを作成
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
    
    param, err := collectVpgCreateVpcPeeringConnectionCmdParams()
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

func collectVpgCreateVpcPeeringConnectionCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForVpgCreateVpcPeeringConnectionCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForVpgCreateVpcPeeringConnectionCmd("/virtual_private_gateways/{vpg_id}/vpc_peering_connections"),
    query: buildQueryForVpgCreateVpcPeeringConnectionCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForVpgCreateVpcPeeringConnectionCmd(path string) string {
  
  
  
  
  
  
  
  
  path = strings.Replace(path, "{" + "vpg_id" + "}", VpgCreateVpcPeeringConnectionCmdVpgId, -1)
  
  
  
  
  
  return path
}

func buildQueryForVpgCreateVpcPeeringConnectionCmd() string {
  result := []string{}
  
  
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


func buildBodyForVpgCreateVpcPeeringConnectionCmd() (string, error) {
  if VpgCreateVpcPeeringConnectionCmdBody != "" {
    if strings.HasPrefix(VpgCreateVpcPeeringConnectionCmdBody, "@") {
      fname := strings.TrimPrefix(VpgCreateVpcPeeringConnectionCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if VpgCreateVpcPeeringConnectionCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return VpgCreateVpcPeeringConnectionCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  
  if VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock != "" {
    result["destinationCidrBlock"] = VpgCreateVpcPeeringConnectionCmdDestinationCidrBlock
  }
  
  
  
  if VpgCreateVpcPeeringConnectionCmdPeerOwnerId != "" {
    result["peerOwnerId"] = VpgCreateVpcPeeringConnectionCmdPeerOwnerId
  }
  
  
  
  if VpgCreateVpcPeeringConnectionCmdPeerVpcId != "" {
    result["peerVpcId"] = VpgCreateVpcPeeringConnectionCmdPeerVpcId
  }
  
  
  
  

  

  

  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

