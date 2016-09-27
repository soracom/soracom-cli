package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(GroupsCmd)
}

// GroupsCmd defines 'groups' subcommand
var GroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: TR("groups.cli.summary"),
	Long:  TR(`groups.cli.description`),
}
