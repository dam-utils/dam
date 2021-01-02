package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var removeRepoCmd = &cobra.Command{
	Use:   "removerepo <id>|<name>",
	Aliases: []string{"rr"},
	Short: "Remove registry specified by name or number.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.RemoveRepo(args[0])
	},
}

func init() {
	removeRepoCmd.Flags().BoolVar(&run.RemoveRepoFlags.Force, "force", false, "Remove registry ignoring tag 'default'.")
}
