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
		flags := []string{"--name","--server","--default","--username","--password"}
		for _, arg := range os.Args {
			for _, flag := range flags {
				//TODO Fix if flags --name='--default'
				if strings.Contains(arg, flag) {
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