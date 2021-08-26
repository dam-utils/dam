package cmd

import (
	"dam/driver/conf/option"
	"dam/driver/logger"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: option.Config.Global.GetProjectName(),
		Short: `Docker Application Manager

Version:
  ` + option.Config.Global.GetProjectVersion(),
		Long: `Docker Application Manager

Version:
  ` + option.Config.Global.GetProjectVersion(),
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

func Init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "x", false, "Enable debug mode")

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
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(saveAppCmd)
	rootCmd.AddCommand(purgeAppCmd)
	rootCmd.AddCommand(infoAppCmd)
}

func setDebugMode() {
	if debug {
		logger.DebugMode = true
	}
}
