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
		setDebugMode()
		run.RemoveApp(args[0])
	},
}

func init() {
	removeAppCmd.Flags().BoolVar(&run.RemoveAppFlags.Force, "force", false, "If deletion fails, force delete from database.")
}