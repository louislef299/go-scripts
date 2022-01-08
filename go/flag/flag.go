package main

import (
	"flag"
	"fmt"
)

func main() {
	interactive := flag.Bool("i", false, "Interactive Mode")
	flag.Parse()

	fmt.Println("Interactive value:", *interactive)
}
