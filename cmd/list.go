package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed your applications.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		run.List()
	},
}

func init() {
	listCmd.Flags().BoolVar(&run.ListFlags.Raw, "raw", false, "List all installed your applications in RAW format.")
}