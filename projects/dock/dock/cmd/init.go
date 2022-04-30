/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
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
		// docker run --name <optional-name> -dt ubuntu
		// docker exec -it <optional-name> "/bin/bash"

		d := exec.Command("docker", "run", "--name", "temp", "-dt", "ubuntu")
		err := d.Run()
		if err != nil {
			log.Fatal(err)
		}

		d = exec.Command("docker", "exec", "-it", "temp", "/bin/bash")
		d.Stderr = os.Stderr
		d.Stdout = os.Stdout
		d.Stdin = os.Stdin

		err = d.Run()
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
