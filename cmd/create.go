package cmd

import (
	"dam/driver/conf/option"
	"dam/run"

	"github.com/spf13/cobra"
)

var createAppCmd = &cobra.Command{
	Use:   "create <project directory>",
	Aliases: []string{"ce"},
	Short: "Create docker application.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		setDebugMode()
		run.CreateApp(args[0])
	},
}

func init(){
	createAppCmd.Flags().StringVar(&run.CreateAppFlags.Name, "name", "", "App name. Sets the value of the "+option.Config.ReservedEnvs.GetAppNameEnv()+" variable.")
	createAppCmd.Flags().StringVar(&run.CreateAppFlags.Version, "version", "", "App version. Sets the value of the "+option.Config.ReservedEnvs.GetAppVersionEnv()+" variable.")
	createAppCmd.Flags().StringVar(&run.CreateAppFlags.Family, "family", "", "App family. Sets the value of the "+option.Config.ReservedEnvs.GetAppFamilyEnv()+" variable.")
	createAppCmd.Flags().BoolVar(&run.CreateAppFlags.MultiVersion, "multiversion", false, "App multiversion. Sets the value of the "+option.Config.ReservedEnvs.GetAppMultiversionEnv()+" variable.")
}
