package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var installAppCmd = &cobra.Command{
	Use:   "install [<app>:<version> | <archive path>]",
	Aliases: []string{"in"},
	Short: "Install docker application from a docker registry or a specific file archive.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.InstallApp(args[0])
	},
}
