package main

func generateTrunkCommands(templateDir, outputDir string) error {
	subCommandTemplate, err := openTemplateFile(templateDir, "trunk.tpl")
	if err != nil {
		return err
	}

	argsSlice := []commandArgs{
		{
			Use:                       "bills",
			Short:                     "bills.cli.summary",
			Long:                      "bills.cli.description",
			CommandVariableName:       "BillsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "bills.go",
		},
		{
			Use:                       "coupons",
			Short:                     "coupons.cli.summary",
			Long:                      "coupons.cli.description",
			CommandVariableName:       "CouponsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "coupons.go",
		},
		{
			Use:                       "credentials",
			Short:                     "credentials.cli.summary",
			Long:                      "credentials.cli.description",
			CommandVariableName:       "CredentialsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "credentials.go",
		},
		{
			Use:                       "data",
			Short:                     "data.cli.summary",
			Long:                      "data.cli.description",
			CommandVariableName:       "DataCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "data.go",
		},
		{
			Use:                       "event-handlers",
			Short:                     "event_handlers.cli.summary",
			Long:                      "event_handlers.cli.description",
			CommandVariableName:       "EventHandlersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "event_handlers.go",
		},
		{
			Use:                       "files",
			Short:                     "files.cli.summary",
			Long:                      "files.cli.description",
			CommandVariableName:       "FilesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "files.go",
		},
		{
			Use:                       "groups",
			Short:                     "groups.cli.summary",
			Long:                      "groups.cli.description",
			CommandVariableName:       "GroupsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "groups.go",
		},
		{
			Use:                       "lora-devices",
			Short:                     "lora_devices.cli.summary",
			Long:                      "lora_devices.cli.description",
			CommandVariableName:       "LoraDevicesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "lora_devices.go",
		},
		{
			Use:                       "lora-gateways",
			Short:                     "lora_gateways.cli.summary",
			Long:                      "lora_gateways.cli.description",
			CommandVariableName:       "LoraGatewaysCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "lora_gateways.go",
		},
		{
			Use:                       "lora-network-sets",
			Short:                     "lora_network_sets.cli.summary",
			Long:                      "lora_network_sets.cli.description",
			CommandVariableName:       "LoraNetworkSetsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "lora_network_sets.go",
		},
		{
			Use:                       "operator",
			Short:                     "operator.cli.summary",
			Long:                      "operator.cli.description",
			CommandVariableName:       "OperatorCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "operator.go",
		},
		{
			Use:                       "auth-keys",
			Short:                     "operator.auth_keys.cli.summary",
			Long:                      "operator.auth_keys.cli.description",
			CommandVariableName:       "OperatorAuthKeysCmd",
			ParentCommandVariableName: "OperatorCmd",
			FileName:                  "operator_auth_keys.go",
		},
		{
			Use:                       "orders",
			Short:                     "orders.cli.summary",
			Long:                      "orders.cli.description",
			CommandVariableName:       "OrdersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "orders.go",
		},
		{
			Use:                       "payment-history",
			Short:                     "payment_history.cli.summary",
			Long:                      "payment_history.cli.description",
			CommandVariableName:       "PaymentHistoryCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payment_history.go",
		},
		{
			Use:                       "payment-methods",
			Short:                     "payment_methods.cli.summary",
			Long:                      "payment_methods.cli.description",
			CommandVariableName:       "PaymentMethodsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payment_methods.go",
		},
		{
			Use:                       "payment-statements",
			Short:                     "payment_statements.cli.summary",
			Long:                      "payment_statements.cli.description",
			CommandVariableName:       "PaymentStatementsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payment_statements.go",
		},
		{
			Use:                       "payer-information",
			Short:                     "payment.payer_information.cli.summary",
			Long:                      "payment.payer_information.cli.description",
			CommandVariableName:       "PayerInformationCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payer_information.go",
		},
		{
			Use:                       "webpay",
			Short:                     "payment_methods.webpay.cli.summary",
			Long:                      "payment_methods.webpay.cli.description",
			CommandVariableName:       "PaymentMethodsWebpayCmd",
			ParentCommandVariableName: "PaymentMethodsCmd",
			FileName:                  "payment_methods_webpay.go",
		},
		{
			Use:                       "products",
			Short:                     "products.cli.summary",
			Long:                      "products.cli.description",
			CommandVariableName:       "ProductsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "products.go",
		},
		{
			Use:                       "roles",
			Short:                     "roles.cli.summary",
			Long:                      "roles.cli.description",
			CommandVariableName:       "RolesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "roles.go",
		},
		{
			Use:                       "shipping-addresses",
			Short:                     "shipping_addresses.cli.summary",
			Long:                      "shipping_addresses.cli.description",
			CommandVariableName:       "ShippingAddressesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "shipping_addresses.go",
		},
		{
			Use:                       "stats",
			Short:                     "stats.cli.summary",
			Long:                      "stats.cli.description",
			CommandVariableName:       "StatsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "stats.go",
		},
		{
			Use:                       "air",
			Short:                     "stats.air.cli.summary",
			Long:                      "stats.air.cli.description",
			CommandVariableName:       "StatsAirCmd",
			ParentCommandVariableName: "StatsCmd",
			FileName:                  "stats_air.go",
		},
		{
			Use:                       "beam",
			Short:                     "stats.beam.cli.summary",
			Long:                      "stats.beam.cli.description",
			CommandVariableName:       "StatsBeamCmd",
			ParentCommandVariableName: "StatsCmd",
			FileName:                  "stats_beam.go",
		},
		{
			Use:                       "subscribers",
			Short:                     "subscribers.cli.summary",
			Long:                      "subscribers.cli.description",
			CommandVariableName:       "SubscribersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "subscribers.go",
		},
		{
			Use:                       "users",
			Short:                     "users.cli.summary",
			Long:                      "users.cli.description",
			CommandVariableName:       "UsersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "users.go",
		},
		{
			Use:                       "auth-keys",
			Short:                     "users.auth_keys.cli.summary",
			Long:                      "users.auth_keys.cli.description",
			CommandVariableName:       "UsersAuthKeysCmd",
			ParentCommandVariableName: "UsersCmd",
			FileName:                  "users_auth_keys.go",
		},
		{
			Use:                       "password",
			Short:                     "users.password.cli.summary",
			Long:                      "users.password.cli.description",
			CommandVariableName:       "UsersPasswordCmd",
			ParentCommandVariableName: "UsersCmd",
			FileName:                  "users_password.go",
		},
		{
			Use:                       "permissions",
			Short:                     "users.permissions.cli.summary",
			Long:                      "users.permissions.cli.description",
			CommandVariableName:       "UsersPermissionsCmd",
			ParentCommandVariableName: "UsersCmd",
			FileName:                  "users_permissions.go",
		},
		{
			Use:                       "vpg",
			Short:                     "virtual_private_gateway.cli.summary",
			Long:                      "virtual_private_gateway.cli.description",
			CommandVariableName:       "VpgCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "vpg.go",
		},
		{
			Use:                       "logs",
			Short:                     "logs.cli.summary",
			Long:                      "logs.cli.description",
			CommandVariableName:       "LogsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "logs.go",
		},
	}

	for _, args := range argsSlice {
		f, err := openOutputFile(outputDir, args.FileName)
		if err != nil {
			return err
		}
		err = subCommandTemplate.Execute(f, args)
		if err != nil {
			return err
		}
	}

	return nil
}
