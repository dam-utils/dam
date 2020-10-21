package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var purgeAppCmd = &cobra.Command{
	Use:   "purge",
	Short: "Remove docker images not attached to apps.",
	Long:  ``,
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.Purge()
	},
}

func init() {
	purgeAppCmd.Flags().BoolVar(&run.PurgeFlags.All, "all", false, "Remove all not used docker images.")
}