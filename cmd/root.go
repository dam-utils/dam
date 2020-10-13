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
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(saveAppCmd)
	rootCmd.AddCommand(purgeAppCmd)
	rootCmd.AddCommand(infoAppCmd)

	//if pFlagDebug && config.DISABLE_DEBUG == false {
	//if debug {
		logger.DebugMode = true
	//}
}
