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
	Short: TR("event_handlers.cli.summary"),
	Long:  TR(`event_handlers.cli.description`),
}
