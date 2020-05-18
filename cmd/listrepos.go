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

var listReposCmd = &cobra.Command{
	Use:   "listrepos (lr) [--raw]",
	Short: "List all defined repositories.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		run.ListRepos()
	},
}

func init() {
	listReposCmd.Flags().BoolVar(&run.ListReposFlags.Raw, "raw", false, "List all defined repositories in RAW format")
}