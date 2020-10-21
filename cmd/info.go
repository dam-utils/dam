package cmd

import (
	"dam/run"
	"github.com/spf13/cobra"
)

var infoAppCmd = &cobra.Command{
	Use:   "info <tag>",
	Short: "Information for your application.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.InfoApp(args[0])
	},
}