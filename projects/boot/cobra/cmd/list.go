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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the contents of the boot directory",
	Long: `Lists the contents of the boot directory.
For example:
	boot list`,
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		CheckError(err)
		files, err := ioutil.ReadDir(fmt.Sprintf("%v/.boot", homeDir))
		CheckError(err)

		if len(files) == 0 {
			fmt.Println("There are no files in boot")
			return
		}

		fmt.Println("The following files are in boot:")
		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
