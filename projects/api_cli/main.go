/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package main

import (
	"os"

	"github.com/louislef299/bash/projects/mlctl/cmd"
)

func main() {
	command := cmd.NewCmdRoot()
	cmd.AddInfraCommands(command)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
