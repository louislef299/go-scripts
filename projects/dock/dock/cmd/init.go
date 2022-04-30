/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var image string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a docker session",
	Long: `This command will spin up an docker container
with an image and exec into it. The default
image is ubuntu`,
	Example: "  dock init",
	Run: func(cmd *cobra.Command, args []string) {
		// docker run --name <optional-name> -dt ubuntu
		// docker exec -it <optional-name> "/bin/bash"

		d := exec.Command("docker", "run", "--name", "temp", "-dt", "ubuntu")
		err := d.Run()
		CheckError(err)

		d = exec.Command("docker", "exec", "-it", "temp", "/bin/bash")
		d.Stderr = os.Stderr
		d.Stdout = os.Stdout
		d.Stdin = os.Stdin

		err = d.Run()
		CheckError(err)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&image, "image", "i", "ubuntu", "The image to use")
}
