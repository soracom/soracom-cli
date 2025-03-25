// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DataDeleteEntryCmdResourceId holds value of 'resource_id' option
var DataDeleteEntryCmdResourceId string

// DataDeleteEntryCmdResourceType holds value of 'resource_type' option
var DataDeleteEntryCmdResourceType string

// DataDeleteEntryCmdTime holds value of 'time' option
var DataDeleteEntryCmdTime int64

func InitDataDeleteEntryCmd() {
	DataDeleteEntryCmd.Flags().StringVar(&DataDeleteEntryCmdResourceId, "resource-id", "", TRAPI("ID of data source resource. The ID to be specified depends on the value of 'resource_type'.| 'resource_type' | The ID you specify ||-|-|| 'Subscriber' | IMSI of the IoT SIM || 'LoraDevice' | ID of the LoRaWAN device || 'Sim' | SIM ID of the IoT SIM || 'SigfoxDevice' | ID of the Sigfox device || 'Device' | ID of the Inventory device || 'SoraCam' | Device ID of the compatible camera device |"))

	DataDeleteEntryCmd.Flags().StringVar(&DataDeleteEntryCmdResourceType, "resource-type", "", TRAPI("Type of data source resource."))

	DataDeleteEntryCmd.Flags().Int64Var(&DataDeleteEntryCmdTime, "time", 0, TRAPI("Timestamp of the target data entry to delete (UNIX time in milliseconds)."))

	DataDeleteEntryCmd.RunE = DataDeleteEntryCmdRunE

	DataCmd.AddCommand(DataDeleteEntryCmd)
}

// DataDeleteEntryCmd defines 'delete-entry' subcommand
var DataDeleteEntryCmd = &cobra.Command{
	Use:   "delete-entry",
	Short: TRAPI("/data/{resource_type}/{resource_id}/{time}:delete:summary"),
	Long:  TRAPI(`/data/{resource_type}/{resource_id}/{time}:delete:description`) + "\n\n" + createLinkToAPIReference("DataEntry", "deleteDataEntry"),
}

func DataDeleteEntryCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectDataDeleteEntryCmdParams(ac)
	if err != nil {
		return err
	}

	body, err := ac.callAPI(param)
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	if body == "" {
		return nil
	}

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectDataDeleteEntryCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("resource_id", "resource-id", "path", parsedBody, DataDeleteEntryCmdResourceId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("resource_type", "resource-type", "path", parsedBody, DataDeleteEntryCmdResourceType)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("time", "time", "path", parsedBody, DataDeleteEntryCmdTime)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDataDeleteEntryCmd("/data/{resource_type}/{resource_id}/{time}"),
		query:  buildQueryForDataDeleteEntryCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDataDeleteEntryCmd(path string) string {

	escapedResourceId := url.PathEscape(DataDeleteEntryCmdResourceId)

	path = strReplace(path, "{"+"resource_id"+"}", escapedResourceId, -1)

	escapedResourceType := url.PathEscape(DataDeleteEntryCmdResourceType)

	path = strReplace(path, "{"+"resource_type"+"}", escapedResourceType, -1)

	path = strReplace(path, "{"+"time"+"}", url.PathEscape(sprintf("%d", DataDeleteEntryCmdTime)), -1)

	return path
}

func buildQueryForDataDeleteEntryCmd() url.Values {
	result := url.Values{}

	return result
}
