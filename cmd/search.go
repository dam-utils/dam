package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [<mask>]",
	Aliases: []string{"se"},
	Short: "Search app in remote registry.",
	Long:  ``,
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			run.Search("")
		} else {
			run.Search(args[0])
		}
	},
}
