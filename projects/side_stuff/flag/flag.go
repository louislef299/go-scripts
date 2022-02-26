package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	interactive := flag.Bool("i", false, "Interactive Mode")
	flag.Parse()

	fmt.Println("Interactive value:", *interactive)

	sum := int(math.Pow(10, 31))
	fmt.Println(sum)

	fmt.Println(43 >> 1)
}
