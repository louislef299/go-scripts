package main

import (
	"flag"
	"fmt"
)

func main() {
	init := flag.String("init", "", "The init flag for docker instantiation")
	help := flag.Bool("help", false, "Show help commands")
	flag.Parse()

	if *help {
		fmt.Println("Usage: fucking [commands]")
		return
	}

	if *init != "" {
		Init(*init)
	}

}

func CheckError(err error) {
	if err != nil {
		fmt.Println("The following error was thrown:", err)
		panic(err)
	}
}
