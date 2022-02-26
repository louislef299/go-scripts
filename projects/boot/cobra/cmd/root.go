/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Move files to/from boot directory",
	Long: `This command will 'boot' a file from its working
directory to the boot directory located in ~/.boot.
For example: 
	* boot <file>
	* boot recover <file>`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: boot <files>")
			return
		}

		workingDir, err := os.Getwd()
		CheckError(err)

		for _, bootFile := range args {
			path, err := os.Getwd()
			CheckError(err)
			exists := DoesFileExist(bootFile, path)
			if !exists {
				fmt.Println("Could not find file", bootFile, "in", workingDir, "to boot!")
				continue
			}

			bootPathFrom := fmt.Sprintf("%v/%v", workingDir, bootFile)
			homeDir, err := os.UserHomeDir()
			CheckError(err)
			bootPathTo := fmt.Sprintf("%v/.boot/%v", homeDir, bootFile)
			err = os.Rename(bootPathFrom, bootPathTo)
			CheckError(err)
			fmt.Println("Booted", bootFile, "to", bootPathTo)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Check if there is an existing command
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err != nil || cmd == nil {
		// Not found
		args := os.Args[1:]
		rootCmd.SetArgs(args)
		rootCmd.Run(rootCmd, args)
	} else {
		err = rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DoesFileExist(bootFile string, path string) bool {
	files, err := ioutil.ReadDir(path)
	CheckError(err)

	fileFound := false
	for _, file := range files {
		if file.Name() == bootFile {
			fileFound = true
		}
	}

	return fileFound
}
