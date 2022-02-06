package main

import (
	"os"
	"os/exec"
)

func Init(name string) {
	commandPath := "./docker_bash/fucking-init.sh"
	cmd := &exec.Cmd{
		Path:   commandPath,
		Args:   []string{commandPath, name},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err := cmd.Start()
	CheckError(err)
}
