package main

import (
	"fmt"

	"github.com/louislef299/comptime-login/pkg/login"
	_ "github.com/louislef299/comptime-login/plugins/aws"
	_ "github.com/louislef299/comptime-login/plugins/docker"
	_ "github.com/louislef299/comptime-login/plugins/k8s"
)

func main() {
	drivers := login.Drivers()

	fmt.Println("found the following drivers:")
	for _, d := range drivers {
		fmt.Println(d)
	}

	if err := login.DLogin("aws"); err != nil {
		panic(err)
	}
	if err := login.DLogin("k8s"); err != nil {
		panic(err)
	}
	if err := login.DLogin("docker"); err != nil {
		panic(err)
	}
}
