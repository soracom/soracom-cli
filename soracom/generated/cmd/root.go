// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd defines 'soracom' command
var RootCmd = &cobra.Command{
	Use:   "soracom",
	Short: "soracom command",
	Long:  `A command line tool to invoke SORACOM API`,
}

var specifiedProfileName string
var specifiedCoverageType string
var providedAPIKey string
var providedAPIToken string
var providedAuthKeyID string
var providedAuthKey string
var providedProfileCommand string
var rawOutput bool
var noRetryOnError bool

func InitRootCmd() {
	RootCmd.PersistentFlags().StringVar(&specifiedProfileName, "profile", "", TRCLI("cli.global-flags.profile"))
	RootCmd.PersistentFlags().StringVar(&specifiedCoverageType, "coverage-type", "", TRCLI("cli.global-flags.coverage-type"))
	RootCmd.PersistentFlags().StringVar(&providedAPIKey, "api-key", "", TRCLI("cli.global-flags.api-key"))
	RootCmd.PersistentFlags().StringVar(&providedAPIToken, "api-token", "", TRCLI("cli.global-flags.api-token"))
	RootCmd.PersistentFlags().StringVar(&providedAuthKeyID, "auth-key-id", "", TRCLI("cli.global-flags.auth-key-id"))
	RootCmd.PersistentFlags().StringVar(&providedAuthKey, "auth-key", "", TRCLI("cli.global-flags.auth-key"))
	RootCmd.PersistentFlags().StringVar(&providedProfileCommand, "profile-command", "", TRCLI("cli.global-flags.profile-command"))
	RootCmd.PersistentFlags().BoolVar(&rawOutput, "raw-output", false, TRCLI("cli.global-flags.raw-output"))
	RootCmd.PersistentFlags().BoolVar(&noRetryOnError, "no-retry-on-error", false, TRCLI("cli.global-flags.no-retry-on-error"))

	InitAllSubCommands()
}

func InitAllSubCommands() {
	InitAuditLogsApiGetCmd()
	InitAuditLogsNapterGetCmd()
	InitAuthCmd()
	InitAuthIssuePasswordResetTokenCmd()
	InitAuthSwitchUserCmd()
	InitAuthVerifyPasswordResetTokenCmd()
	InitBillsExportCmd()
	InitBillsExportLatestCmd()
	InitBillsGetCmd()
	InitBillsGetDailyCmd()
	InitBillsGetLatestCmd()
	InitBillsListCmd()
	InitBillsSummariesGetBillItemsCmd()
	InitBillsSummariesGetSimsCmd()
	InitCellLocationsBatchGetCmd()
	InitCellLocationsGetCmd()
	InitCouponsConfirmCmd()
	InitCouponsCreateCmd()
	InitCouponsListCmd()
	InitCouponsRegisterCmd()
	InitCredentialsCreateCmd()
	InitCredentialsDeleteCmd()
	InitCredentialsListCmd()
	InitCredentialsUpdateCmd()
	InitDataDeleteEntryCmd()
	InitDataGetCmd()
	InitDataGetEntriesCmd()
	InitDataGetEntryCmd()
	InitDataListSourceResourcesCmd()
	InitDevicesCreateCmd()
	InitDevicesCreateObjectModelCmd()
	InitDevicesDeleteCmd()
	InitDevicesDeleteDeviceTagCmd()
	InitDevicesDeleteObjectModelCmd()
	InitDevicesExecuteResourceCmd()
	InitDevicesGetCmd()
	InitDevicesGetDataCmd()
	InitDevicesGetInstanceCmd()
	InitDevicesGetObjectModelCmd()
	InitDevicesGetResourceCmd()
	InitDevicesListCmd()
	InitDevicesListObjectModelsCmd()
	InitDevicesObserveResourceCmd()
	InitDevicesObserveResourcesCmd()
	InitDevicesPutDeviceTagsCmd()
	InitDevicesPutResourceCmd()
	InitDevicesSetGroupCmd()
	InitDevicesSetObjectModelScopeCmd()
	InitDevicesUnobserveResourceCmd()
	InitDevicesUnobserveResourcesCmd()
	InitDevicesUnsetGroupCmd()
	InitDevicesUpdateObjectModelCmd()
	InitDiagnosticsGetCmd()
	InitDiagnosticsSendRequestCmd()
	InitEmailsDeleteCmd()
	InitEmailsGetCmd()
	InitEmailsIssueAddEmailTokenCmd()
	InitEmailsListCmd()
	InitEmailsVerifyAddEmailTokenCmd()
	InitEventHandlersCreateCmd()
	InitEventHandlersDeleteCmd()
	InitEventHandlersGetCmd()
	InitEventHandlersIgnoreCmd()
	InitEventHandlersListCmd()
	InitEventHandlersListForSubscriberCmd()
	InitEventHandlersUnignoreCmd()
	InitEventHandlersUpdateCmd()
	InitFilesDeleteCmd()
	InitFilesDeleteDirectoryCmd()
	InitFilesFindCmd()
	InitFilesGetCmd()
	InitFilesGetExportedCmd()
	InitFilesGetMetadataCmd()
	InitFilesListCmd()
	InitFilesPutCmd()
	InitGadgetsDeleteTagCmd()
	InitGadgetsDisableTerminationCmd()
	InitGadgetsEnableTerminationCmd()
	InitGadgetsGetCmd()
	InitGadgetsListCmd()
	InitGadgetsPutTagsCmd()
	InitGadgetsRegisterCmd()
	InitGadgetsTerminateCmd()
	InitGroupsCreateCmd()
	InitGroupsDeleteCmd()
	InitGroupsDeleteConfigCmd()
	InitGroupsDeleteConfigNamespaceCmd()
	InitGroupsDeleteTagCmd()
	InitGroupsGetCmd()
	InitGroupsListCmd()
	InitGroupsListSubscribersCmd()
	InitGroupsPutConfigCmd()
	InitGroupsPutTagsCmd()
	InitLagoonCreateUserCmd()
	InitLagoonDeleteUserCmd()
	InitLagoonGetImageLinkCmd()
	InitLagoonLicensePacksListStatusCmd()
	InitLagoonLicensePacksUpdateCmd()
	InitLagoonListUsersCmd()
	InitLagoonRegisterCmd()
	InitLagoonTerminateCmd()
	InitLagoonUpdateUserEmailCmd()
	InitLagoonUpdateUserPasswordCmd()
	InitLagoonUpdateUserPermissionCmd()
	InitLagoonUpdatedPlanCmd()
	InitLagoonUsersCreateCmd()
	InitLagoonUsersDeleteCmd()
	InitLagoonUsersListCmd()
	InitLagoonUsersUpdateEmailCmd()
	InitLagoonUsersUpdatePasswordCmd()
	InitLagoonUsersUpdatePermissionCmd()
	InitLogoutCmd()
	InitLogsGetCmd()
	InitLoraDevicesDeleteTagCmd()
	InitLoraDevicesDisableTerminationCmd()
	InitLoraDevicesEnableTerminationCmd()
	InitLoraDevicesGetCmd()
	InitLoraDevicesGetDataCmd()
	InitLoraDevicesListCmd()
	InitLoraDevicesPutTagsCmd()
	InitLoraDevicesRegisterCmd()
	InitLoraDevicesSendDataCmd()
	InitLoraDevicesSetGroupCmd()
	InitLoraDevicesTerminateCmd()
	InitLoraDevicesUnsetGroupCmd()
	InitLoraGatewaysDeleteTagCmd()
	InitLoraGatewaysDisableTerminationCmd()
	InitLoraGatewaysEnableTerminationCmd()
	InitLoraGatewaysGetCmd()
	InitLoraGatewaysListCmd()
	InitLoraGatewaysPutTagsCmd()
	InitLoraGatewaysSetNetworkSetCmd()
	InitLoraGatewaysTerminateCmd()
	InitLoraGatewaysUnsetNetworkSetCmd()
	InitLoraNetworkSetsAddPermissionCmd()
	InitLoraNetworkSetsCreateCmd()
	InitLoraNetworkSetsDeleteCmd()
	InitLoraNetworkSetsDeleteTagCmd()
	InitLoraNetworkSetsGetCmd()
	InitLoraNetworkSetsListCmd()
	InitLoraNetworkSetsListGatewaysCmd()
	InitLoraNetworkSetsPutTagCmd()
	InitLoraNetworkSetsRevokePermissionCmd()
	InitOperatorAddContractCmd()
	InitOperatorAddCoverageTypeCmd()
	InitOperatorAuthKeysDeleteCmd()
	InitOperatorAuthKeysGenerateCmd()
	InitOperatorAuthKeysListCmd()
	InitOperatorConfigurationDeleteCmd()
	InitOperatorConfigurationGetCmd()
	InitOperatorConfigurationSetCmd()
	InitOperatorCreateCompanyInformationCmd()
	InitOperatorCreateIndividualInformationCmd()
	InitOperatorDeleteContractCmd()
	InitOperatorEnableMfaCmd()
	InitOperatorGenerateApiTokenCmd()
	InitOperatorGetCmd()
	InitOperatorGetCompanyInformationCmd()
	InitOperatorGetIndividualInformationCmd()
	InitOperatorGetMfaStatusCmd()
	InitOperatorGetSupportTokenCmd()
	InitOperatorIssueMfaRevokeTokenCmd()
	InitOperatorRevokeMfaCmd()
	InitOperatorRevokeOperatorAuthTokensCmd()
	InitOperatorUpdateCompanyInformationCmd()
	InitOperatorUpdateIndividualInformationCmd()
	InitOperatorUpdatePasswordCmd()
	InitOperatorVerifyMfaOtpCmd()
	InitOperatorVerifyMfaRevokeTokenCmd()
	InitOrdersCancelCmd()
	InitOrdersConfirmCmd()
	InitOrdersCreateCmd()
	InitOrdersGetCmd()
	InitOrdersListCmd()
	InitOrdersListSubscribersCmd()
	InitOrdersRegisterSubscribersCmd()
	InitOrdersResourceInitialSettingUpdateCmd()
	InitPayerInformationGetCmd()
	InitPayerInformationRegisterCmd()
	InitPaymentHistoryGetCmd()
	InitPaymentMethodsGetCurrentCmd()
	InitPaymentMethodsReactivateCurrentCmd()
	InitPaymentStatementsExportCmd()
	InitPaymentStatementsListCmd()
	InitPortMappingsCreateCmd()
	InitPortMappingsDeleteCmd()
	InitPortMappingsGetCmd()
	InitPortMappingsListCmd()
	InitPortMappingsListForSimCmd()
	InitPortMappingsListForSubscriberCmd()
	InitProductsListCmd()
	InitQueryDevicesCmd()
	InitQuerySigfoxDevicesCmd()
	InitQuerySimsCmd()
	InitQuerySubscribersCmd()
	InitQuerySubscribersTrafficVolumeRankingCmd()
	InitQueryTrafficRankingCmd()
	InitResourceSummariesGetCmd()
	InitRolesCreateCmd()
	InitRolesDeleteCmd()
	InitRolesGetCmd()
	InitRolesListCmd()
	InitRolesListUsersCmd()
	InitRolesUpdateCmd()
	InitSandboxCouponsCreateCmd()
	InitSandboxInitCmd()
	InitSandboxOperatorsDeleteCmd()
	InitSandboxOperatorsGetSignupTokenCmd()
	InitSandboxOrdersShipCmd()
	InitSandboxStatsAirInsertCmd()
	InitSandboxStatsBeamInsertCmd()
	InitSandboxSubscribersCreateCmd()
	InitShippingAddressesCreateCmd()
	InitShippingAddressesDeleteCmd()
	InitShippingAddressesGetCmd()
	InitShippingAddressesListCmd()
	InitShippingAddressesUpdateCmd()
	InitSigfoxDevicesDeleteTagCmd()
	InitSigfoxDevicesDisableTerminationCmd()
	InitSigfoxDevicesEnableTerminationCmd()
	InitSigfoxDevicesGetCmd()
	InitSigfoxDevicesGetDataCmd()
	InitSigfoxDevicesListCmd()
	InitSigfoxDevicesPutTagsCmd()
	InitSigfoxDevicesRegisterCmd()
	InitSigfoxDevicesSendDataCmd()
	InitSigfoxDevicesSetGroupCmd()
	InitSigfoxDevicesTerminateCmd()
	InitSigfoxDevicesUnsetGroupCmd()
	InitSimProfileOrdersConfirmCmd()
	InitSimProfileOrdersCreateCmd()
	InitSimProfileOrdersDeleteCmd()
	InitSimProfileOrdersGetCmd()
	InitSimProfileOrdersListCmd()
	InitSimProfileOrdersListProfilesCmd()
	InitSimsActivateCmd()
	InitSimsAddSubscriptionCmd()
	InitSimsAttachArcCredentialsCmd()
	InitSimsCancelSubscriptionContainerDownloadCmd()
	InitSimsCreateArcSessionCmd()
	InitSimsCreateCmd()
	InitSimsCreatePacketCaptureSessionCmd()
	InitSimsDeactivateCmd()
	InitSimsDeletePacketCaptureSessionCmd()
	InitSimsDeleteSessionCmd()
	InitSimsDeleteTagCmd()
	InitSimsDisableTerminationCmd()
	InitSimsDownlinkHttpCmd()
	InitSimsDownlinkPingCmd()
	InitSimsDownlinkSshCmd()
	InitSimsEnableTerminationCmd()
	InitSimsGetCmd()
	InitSimsGetDataCmd()
	InitSimsGetPacketCaptureSessionCmd()
	InitSimsListCmd()
	InitSimsListPacketCaptureSessionsCmd()
	InitSimsListStatusHistoryCmd()
	InitSimsPutTagsCmd()
	InitSimsRegisterCmd()
	InitSimsRemoveArcCredentialsCmd()
	InitSimsRenewArcCredentialsCmd()
	InitSimsReportLocalInfoCmd()
	InitSimsSendSmsCmd()
	InitSimsSessionEventsCmd()
	InitSimsSetExpiryTimeCmd()
	InitSimsSetGroupCmd()
	InitSimsSetImeiLockCmd()
	InitSimsSetToStandbyCmd()
	InitSimsStopPacketCaptureSessionCmd()
	InitSimsSuspendCmd()
	InitSimsTerminateCmd()
	InitSimsTerminateSubscriptionContainerCmd()
	InitSimsUnsetExpiryTimeCmd()
	InitSimsUnsetGroupCmd()
	InitSimsUnsetImeiLockCmd()
	InitSimsUpdateSpeedClassCmd()
	InitSoraCamDevicesAtomCamAssignLicenseCmd()
	InitSoraCamDevicesAtomCamGetFirmwareUpdateCmd()
	InitSoraCamDevicesAtomCamGetSettingsCmd()
	InitSoraCamDevicesAtomCamListFirmwareUpdatesCmd()
	InitSoraCamDevicesAtomCamSettingsGetLogoCmd()
	InitSoraCamDevicesAtomCamSettingsGetMotionCmd()
	InitSoraCamDevicesAtomCamSettingsGetMotionSensitivityCmd()
	InitSoraCamDevicesAtomCamSettingsGetMotionTaggingCmd()
	InitSoraCamDevicesAtomCamSettingsGetNightVisionCmd()
	InitSoraCamDevicesAtomCamSettingsGetQualityCmd()
	InitSoraCamDevicesAtomCamSettingsGetRotationCmd()
	InitSoraCamDevicesAtomCamSettingsGetSoundCmd()
	InitSoraCamDevicesAtomCamSettingsGetSoundSensitivityCmd()
	InitSoraCamDevicesAtomCamSettingsGetStatusLightCmd()
	InitSoraCamDevicesAtomCamSettingsGetTimestampCmd()
	InitSoraCamDevicesAtomCamSettingsSetLogoCmd()
	InitSoraCamDevicesAtomCamSettingsSetMotionCmd()
	InitSoraCamDevicesAtomCamSettingsSetMotionSensitivityCmd()
	InitSoraCamDevicesAtomCamSettingsSetMotionTaggingCmd()
	InitSoraCamDevicesAtomCamSettingsSetNightVisionCmd()
	InitSoraCamDevicesAtomCamSettingsSetQualityCmd()
	InitSoraCamDevicesAtomCamSettingsSetRotationCmd()
	InitSoraCamDevicesAtomCamSettingsSetSoundCmd()
	InitSoraCamDevicesAtomCamSettingsSetSoundSensitivityCmd()
	InitSoraCamDevicesAtomCamSettingsSetStatusLightCmd()
	InitSoraCamDevicesAtomCamSettingsSetTimestampCmd()
	InitSoraCamDevicesAtomCamUnassignLicenseCmd()
	InitSoraCamDevicesAtomCamUpdateFirmwareCmd()
	InitSoraCamDevicesDataCreateEntryCmd()
	InitSoraCamDevicesDataUpdateEntryCmd()
	InitSoraCamDevicesDeleteCmd()
	InitSoraCamDevicesDeleteExportUsageLimitOverrideCmd()
	InitSoraCamDevicesEventsListCmd()
	InitSoraCamDevicesEventsListForDeviceCmd()
	InitSoraCamDevicesGetCmd()
	InitSoraCamDevicesGetExportUsageCmd()
	InitSoraCamDevicesGetNameCmd()
	InitSoraCamDevicesGetPowerStateCmd()
	InitSoraCamDevicesGetStreamingVideoCmd()
	InitSoraCamDevicesImagesExportCmd()
	InitSoraCamDevicesImagesGetExportedCmd()
	InitSoraCamDevicesImagesListExportsCmd()
	InitSoraCamDevicesImagesListExportsForDeviceCmd()
	InitSoraCamDevicesListCmd()
	InitSoraCamDevicesRebootCmd()
	InitSoraCamDevicesRecordingsAndEventsListForDeviceCmd()
	InitSoraCamDevicesSetNameCmd()
	InitSoraCamDevicesSetPowerStateCmd()
	InitSoraCamDevicesUpdateExportUsageLimitOverrideCmd()
	InitSoraCamDevicesVideosExportCmd()
	InitSoraCamDevicesVideosGetExportedCmd()
	InitSoraCamDevicesVideosListExportsCmd()
	InitSoraCamDevicesVideosListExportsForDeviceCmd()
	InitSoraCamLicensePacksListCmd()
	InitSoraCamLicensePacksUpdateQuantityCmd()
	InitSoraletsCreateCmd()
	InitSoraletsDeleteCmd()
	InitSoraletsDeleteVersionCmd()
	InitSoraletsExecCmd()
	InitSoraletsGetCmd()
	InitSoraletsGetLogsCmd()
	InitSoraletsListCmd()
	InitSoraletsListVersionsCmd()
	InitSoraletsTestCmd()
	InitSoraletsUploadCmd()
	InitStatsAirExportCmd()
	InitStatsAirGetCmd()
	InitStatsAirGroupsGetCmd()
	InitStatsAirOperatorsGetCmd()
	InitStatsAirSimsGetCmd()
	InitStatsBeamExportCmd()
	InitStatsBeamGetCmd()
	InitStatsFunkExportCmd()
	InitStatsFunkGetCmd()
	InitStatsFunnelExportCmd()
	InitStatsFunnelGetCmd()
	InitStatsHarvestExportCmd()
	InitStatsHarvestOperatorsGetCmd()
	InitStatsHarvestSubscribersGetCmd()
	InitStatsNapterAuditLogsGetCmd()
	InitSubscribersActivateCmd()
	InitSubscribersDeactivateCmd()
	InitSubscribersDeleteSessionCmd()
	InitSubscribersDeleteTagCmd()
	InitSubscribersDeleteTransferTokenCmd()
	InitSubscribersDisableTerminationCmd()
	InitSubscribersDownlinkPingCmd()
	InitSubscribersEnableTerminationCmd()
	InitSubscribersExportCmd()
	InitSubscribersGetCmd()
	InitSubscribersGetDataCmd()
	InitSubscribersIssueTransferTokenCmd()
	InitSubscribersListCmd()
	InitSubscribersPutBundlesCmd()
	InitSubscribersPutTagsCmd()
	InitSubscribersRegisterCmd()
	InitSubscribersReportLocalInfoCmd()
	InitSubscribersSendSmsByMsisdnCmd()
	InitSubscribersSendSmsCmd()
	InitSubscribersSessionEventsCmd()
	InitSubscribersSetExpiryTimeCmd()
	InitSubscribersSetGroupCmd()
	InitSubscribersSetImeiLockCmd()
	InitSubscribersSetToStandbyCmd()
	InitSubscribersSuspendCmd()
	InitSubscribersTerminateCmd()
	InitSubscribersUnsetExpiryTimeCmd()
	InitSubscribersUnsetGroupCmd()
	InitSubscribersUnsetImeiLockCmd()
	InitSubscribersUpdateSpeedClassCmd()
	InitSubscribersVerifyTransferTokenCmd()
	InitSystemNotificationsDeleteCmd()
	InitSystemNotificationsGetCmd()
	InitSystemNotificationsListCmd()
	InitSystemNotificationsSetCmd()
	InitUsersAttachRoleCmd()
	InitUsersAuthKeysDeleteCmd()
	InitUsersAuthKeysGenerateCmd()
	InitUsersAuthKeysGetCmd()
	InitUsersAuthKeysListCmd()
	InitUsersCreateCmd()
	InitUsersDefaultPermissionsDeleteCmd()
	InitUsersDefaultPermissionsGetCmd()
	InitUsersDefaultPermissionsUpdateCmd()
	InitUsersDeleteCmd()
	InitUsersDetachRoleCmd()
	InitUsersGetCmd()
	InitUsersListCmd()
	InitUsersListRolesCmd()
	InitUsersMfaEnableCmd()
	InitUsersMfaGetCmd()
	InitUsersMfaRevokeCmd()
	InitUsersMfaVerifyCmd()
	InitUsersPasswordConfiguredCmd()
	InitUsersPasswordCreateCmd()
	InitUsersPasswordDeleteCmd()
	InitUsersPasswordUpdateCmd()
	InitUsersPermissionsDeleteCmd()
	InitUsersPermissionsGetCmd()
	InitUsersPermissionsUpdateCmd()
	InitUsersRevokeUserAuthTokensCmd()
	InitUsersTrustPolicyDeleteCmd()
	InitUsersTrustPolicyGetCmd()
	InitUsersTrustPolicyUpdateCmd()
	InitUsersUpdateCmd()
	InitVolumeDiscountsAvailableDiscountsCmd()
	InitVolumeDiscountsConfirmCmd()
	InitVolumeDiscountsCreateCmd()
	InitVolumeDiscountsGetCmd()
	InitVolumeDiscountsListCmd()
	InitVpgAcceptTransitGatewayVpcAttachmentCmd()
	InitVpgCloseGateCmd()
	InitVpgCreateCmd()
	InitVpgCreateCustomerRouteCmd()
	InitVpgCreateMirroringPeerCmd()
	InitVpgCreatePacketCaptureSessionCmd()
	InitVpgCreateTransitGatewayPeeringConnectionCmd()
	InitVpgCreateTransitGatewayVpcAttachmentCmd()
	InitVpgCreateVpcPeeringConnectionCmd()
	InitVpgDeleteCustomerRouteCmd()
	InitVpgDeleteIpAddressMapEntryCmd()
	InitVpgDeleteMirroringPeerCmd()
	InitVpgDeletePacketCaptureSessionCmd()
	InitVpgDeleteTagCmd()
	InitVpgDeleteTransitGatewayPeeringConnectionCmd()
	InitVpgDeleteTransitGatewayVpcAttachmentCmd()
	InitVpgDeleteVpcPeeringConnectionCmd()
	InitVpgDisableGatePrivacySeparatorCmd()
	InitVpgDisableSimBasedRoutingCmd()
	InitVpgEnableGatePrivacySeparatorCmd()
	InitVpgEnableSimBasedRoutingCmd()
	InitVpgGetCmd()
	InitVpgGetPacketCaptureSessionCmd()
	InitVpgListCmd()
	InitVpgListGatePeersCmd()
	InitVpgListIpAddressMapEntriesCmd()
	InitVpgListPacketCaptureSessionsCmd()
	InitVpgOpenGateCmd()
	InitVpgPutIpAddressMapEntryCmd()
	InitVpgPutSimBasedRoutingRoutesCmd()
	InitVpgRegisterGatePeerCmd()
	InitVpgReleaseFixedPublicIpAddressesCmd()
	InitVpgSetFixedPublicIpAddressesCmd()
	InitVpgSetInspectionCmd()
	InitVpgSetRedirectionCmd()
	InitVpgSetRoutingFilterCmd()
	InitVpgSetVxlanIdCmd()
	InitVpgStopPacketCaptureSessionCmd()
	InitVpgTerminateCmd()
	InitVpgUnregisterGatePeerCmd()
	InitVpgUnsetInspectionCmd()
	InitVpgUnsetRedirectionCmd()
	InitVpgUpdateCustomerRouteCmd()
	InitVpgUpdateMirroringPeerCmd()
	InitVpgUpdateTagsCmd()
}
