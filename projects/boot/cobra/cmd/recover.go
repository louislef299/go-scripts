/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.edu>

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
	Short: "Recover files boot directory into current working directory",
	Long: `Recover files boot directory into current working directory.
	For example:
		boot recover example1.txt example2.txt`,
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
}
