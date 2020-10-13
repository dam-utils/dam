package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export <file path>",
	Aliases: []string{"ex"},
	Short: "Export apps to file.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		run.Export(args[0])
	},
}

func init() {
	exportCmd.Flags().BoolVar(&run.ExportFlags.All, "all", false, "Export apps to archive.")
}
