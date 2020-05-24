// Copyright 2020 The Docker Applications Manager Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
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