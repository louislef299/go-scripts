package cmd

import "github.com/spf13/cobra"

func AddInfraCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initCmd)
}
