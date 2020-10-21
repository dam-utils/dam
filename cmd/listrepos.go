package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var listReposCmd = &cobra.Command{
	Use:   "listrepos",
	Aliases: []string{"lr"},
	Short: "List all defined repositories.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.ListRepos()
	},
}

func init() {
	listReposCmd.Flags().BoolVar(&run.ListReposFlags.Raw, "raw", false, "List all defined repositories in RAW format.")
}