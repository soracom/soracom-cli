package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	LagoonCmd.AddCommand(LagoonDashboardsCmd)
}

// LagoonDashboardsCmd defines 'dashboards' subcommand
var LagoonDashboardsCmd = &cobra.Command{
	Use:   "dashboards",
	Short: TRCLI("cli.lagoon.dashboards.summary"),
	Long:  TRCLI(`cli.lagoon.dashboards.description`),
}
