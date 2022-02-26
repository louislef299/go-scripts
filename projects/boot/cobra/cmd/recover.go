/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// recoverCmd represents the recover command
var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Aliases: []string{"rec"},
	Run: func(cmd *cobra.Command, args []string) {
		workingDir, err := os.Getwd()
		CheckError(err)

		for _, bootFile := range args {
			homeDir, err := os.UserHomeDir()
			CheckError(err)
			bootPath := fmt.Sprintf("%v/.boot", homeDir)
			exists := DoesFileExist(bootFile, bootPath)
			if !exists {
				fmt.Println("Could not find file", bootFile, "in", bootPath, "to recover!")
				continue
			}

			bootPathFrom := fmt.Sprintf("%v/.boot/%v", homeDir, bootFile)
			bootPathTo := fmt.Sprintf("%v/%v", workingDir, bootFile)
			err = os.Rename(bootPathFrom, bootPathTo)
			CheckError(err)
			fmt.Println("Recovered", bootFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(recoverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recoverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recoverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
