/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>
*/
package main

import "github.com/louislef299/go-scripts/projects/mlctl/docker"

/*func main() {
	command := cmd.NewCmdRoot()
	cmd.AddInfraCommands(command)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}*/

func main() {
	docker.Init()
}
