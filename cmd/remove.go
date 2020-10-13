package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var removeAppCmd = &cobra.Command{
	Use:   "remove <app>",
	Aliases: []string{"rm"},
	Short: "Remove docker application.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		run.RemoveApp(args[0])
	},
}