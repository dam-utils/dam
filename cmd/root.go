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
		Use:   config.PROJECT_NAME,
		Short: `Docker Application Manager

Version:
  `+config.PROJECT_VERSION,
		Long:  `Docker Application Manager

Version:
  `+config.PROJECT_VERSION,
	}
)

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Fatal("Internal error. Cannot execute root command")
	}
}

var debug bool

func init() {
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")

	rootCmd.AddCommand(removeAppCmd)
	rootCmd.AddCommand(installAppCmd)
	rootCmd.AddCommand(createAppCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(listReposCmd)
	rootCmd.AddCommand(addRepoCmd)
	rootCmd.AddCommand(removeRepoCmd)
	rootCmd.AddCommand(modifyRepoCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(exportCmd)

	//if pFlagDebug && config.DISABLE_DEBUG == false {
	if debug {
		logger.DebugMode = true
	}
}
