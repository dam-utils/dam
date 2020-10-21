package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import <file path>",
	Aliases: []string{"im"},
	Short: "Import apps from file.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.Import(args[0])
	},
}

func init() {
	importCmd.Flags().BoolVar(&run.ImportFlags.Yes, "yes", false, "Agree to all questions.")
	importCmd.Flags().BoolVar(&run.ImportFlags.Restore, "restore", false, "Remove all applications before importing.")
}

