package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Used to delete git branches given a file of branches to delete
// File cannot have extra spaces or anything

// Usage: got -rm=<file>

func main() {
	rm := flag.String("rm", "None", "File to use to remove git branches")
	flag.Parse()

	if *rm == "None" {
	   fmt.Println("Usage: got -rm=<file>")
	   return
	}

	file, err := os.Open(*rm)
	CheckErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		cmd := exec.Command("git", "branch", "-D", scanner.Text())
		stdout, err := cmd.Output()
		// Print the output
		fmt.Println(string(stdout))
		CheckErr(err)
	}

	err = scanner.Err()
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
