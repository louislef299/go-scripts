package main

import (
	"fmt"

	"github.com/louislef299/comptime-login/pkg/login"
	_ "github.com/louislef299/comptime-login/plugins/aws"
)

func main() {
	drivers := login.Drivers()

	fmt.Println("found the following drivers:")
	for _, d := range drivers {
		fmt.Println(d)
	}

	_, err := login.DLogin("aws")
	if err != nil {
		panic(err)
	}
}
