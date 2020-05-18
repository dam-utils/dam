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
	"dam/config"
	"dam/driver/logger"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   config.UTIL_NAME,
		Short: "/--/",
		Long:  `/--/`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(installAppCmd)
	rootCmd.AddCommand(createAppCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(listReposCmd)
	rootCmd.AddCommand(addRepoCmd)
	rootCmd.AddCommand(removeRepoCmd)
	rootCmd.AddCommand(modifyRepoCmd)
	rootCmd.AddCommand(searchCmd)

	//if pFlagDebug && config.DISABLE_DEBUG == false {
	if !config.DISABLE_DEBUG {
		logger.DebugMode = true
	}
}
