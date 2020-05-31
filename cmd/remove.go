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

var removeAppCmd = &cobra.Command{
	Use:   "remove <app>",
	Aliases: []string{"rm"},
	Short: "Remove docker application.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		run.RemoveApp(args[0])
	},
}

func init() {
	// TODO "-f" flag
	//removeAppCmd.Flags().BoolVar(&run.RemoveAppFlags.Force, "force", false, "Removing applications from the database with ignoring errors.")
}