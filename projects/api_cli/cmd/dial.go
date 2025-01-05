/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>
*/
package cmd

import (
	"github.com/louislef299/go-scripts/projects/mlctl/client"
	"github.com/spf13/cobra"
)

// dialCmd represents the dial command
var dialCmd = &cobra.Command{
	Use:   "dial",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := client.Client{Log: Log}
		c.Dial()
	},
}
