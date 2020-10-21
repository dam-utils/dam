package cmd

import (
	"os"
	"strings"

	"dam/run"

	"github.com/spf13/cobra"
)

var modifyRepoCmd = &cobra.Command{
	Use:   "modifyrepo <id>",
	Aliases: []string{"mr"},
	Short: "Modify properties of repositories specified.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		flags := []string{"--name","--server","--default","--username","--password"}
		for _, arg := range os.Args {
			for _, flag := range flags {
				if strings.HasPrefix(arg, flag) {
					run.ExistingMRFlags[flag] = true
				}
			}
		}
		run.ModifyRepo(args[0])
	},
}

func init() {
	modifyRepoCmd.Flags().BoolVar(&run.ModifyRepoFlags.Default, "default", true, "Mark the registry as default.")
	modifyRepoCmd.Flags().StringVar(&run.ModifyRepoFlags.Name, "name", "", "New name of the registry.")
	modifyRepoCmd.Flags().StringVar(&run.ModifyRepoFlags.Server, "server", "", "New IP-address or new hostname of the registry.")
	modifyRepoCmd.Flags().StringVar(&run.ModifyRepoFlags.Username, "username", "", "New username to logger on the remote registry.")
	modifyRepoCmd.Flags().StringVar(&run.ModifyRepoFlags.Password, "password", "", "New password to logger on the remote registry.")
}