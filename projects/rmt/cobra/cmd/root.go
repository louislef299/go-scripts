/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rmt",
	Short: "Removes all files containing a tilda from the current working directory",
	Long: `Removes all files containing a tilda from the current working directory.
	Usage: rmt [ -i | -help ]`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		files := ListDir()
		count := 0
		for _, file := range files {
			if strings.Contains(file.Name(), "~") {
				count++
				fmt.Printf("Removing %v...\n", file.Name())
				err := os.Remove(file.Name())
				CheckError(err)
			}
		}

		if count == 0 {
			fmt.Println("No tilda files found!")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rmt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListDir() []fs.FileInfo {
	path, err := os.Getwd()
	CheckError(err)

	files, err := ioutil.ReadDir(path)
	CheckError(err)

	return files
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
