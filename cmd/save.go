package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var saveAppCmd = &cobra.Command{
	Use:   "save [<repos>/]<app>:<version>",
	Short: "Save app to an archive.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.Save(args[0])
	},
}

func init() {
	saveAppCmd.Flags().StringVar(&run.SaveFlags.FilePath, "f", "", "Set file name for saving archive.")
}