/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("user.name", "Louis")
		viper.WriteConfig()
		Log.Trace.Println("Trace")
		Log.Info.Println("Info")
		Log.Warning.Println("Warning")
		Log.Error.Println("Error")
	},
}
