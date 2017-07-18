package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(EventHandlersCmd)
}

// EventHandlersCmd defines 'event-handlers' subcommand
var EventHandlersCmd = &cobra.Command{
	Use:   "event-handlers",
	Short: TRCLI("cli.event-handlers.summary"),
	Long:  TRCLI(`cli.event-handlers.description`),
}
