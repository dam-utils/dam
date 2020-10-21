package cmd

import (
	"dam/run"

	"github.com/spf13/cobra"
)

var addRepoCmd = &cobra.Command{
	Use:   "addrepo --name <name> --server <hostname>",
	Aliases: []string{"ar"},
	Short: "Add an app registry to the system.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.AddRepo()
	},
}

func init() {
	addRepoCmd.PersistentFlags().StringVar(&run.AddRepoFlags.Name, "name", "", "Repository name. Use only to manage this registry.")
	addRepoCmd.PersistentFlags().StringVar(&run.AddRepoFlags.Server, "server", "", "IP-address or hostname of the remote registry.")
	addRepoCmd.Flags().BoolVarP(&run.AddRepoFlags.Default, "default", "d", false, "Use this registry as default registry. It means that you can ommit registry name in an install command.")
	addRepoCmd.Flags().StringVar(&run.AddRepoFlags.Username, "username", "", "Username to logger on the remote registry.")
	addRepoCmd.Flags().StringVar(&run.AddRepoFlags.Password, "password", "", "Password to logger on the remote registry.")
}