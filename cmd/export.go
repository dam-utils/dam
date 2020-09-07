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
