/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.edu>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Recover all files from boot into current working directory",
	Long: `Recover all files from boot into current working directory. 
For example: 
	boot recover all`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		CheckError(err)
		files, err := ioutil.ReadDir(fmt.Sprintf("%v/.boot", homeDir))
		CheckError(err)
		workingDir, err := os.Getwd()
		CheckError(err)

		for _, file := range files {
			bootPathFrom := fmt.Sprintf("%v/.boot/%v", homeDir, file.Name())
			bootPathTo := fmt.Sprintf("%v/%v", workingDir, file.Name())
			err := os.Rename(bootPathFrom, bootPathTo)
			CheckError(err)
			fmt.Println("Recovered", file.Name())
		}
	},
}

func init() {
	recoverCmd.AddCommand(allCmd)
}
