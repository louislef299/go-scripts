/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"github.com/louislef299/bash/projects/mlctl/server"
	"github.com/spf13/cobra"
)

var port int

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
		s := server.Server{Log: Log}
		s.StartServer(port)
	},
}

func buildInit() {
	initCmd.Flags().IntVarP(&port, "port", "p", 50051, "the port to run the service on")
}
