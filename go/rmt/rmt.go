package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
 * Usage: rmt
 * Function: Removes all files with tildas
 * Options:
 * -i Interactive mode
 */

func main() {
	interactive := flag.Bool("i", false, "Interactive Mode")
	flag.Parse()

	files := ListDir()

	if *interactive {
		InterRemoveFiles(files)
	} else {
		RemoveFiles(files)
	}
}

func ListDir() []fs.FileInfo {
	path, err := os.Getwd()
	CheckError(err)

	files, err := ioutil.ReadDir(path)
	CheckError(err)

	return files
}

func InterRemoveFiles(files []fs.FileInfo) {
	count := 0
	for _, file := range files {
		if strings.Contains(file.Name(), "~") {
			fmt.Print("Would you like to delete ", file.Name(), ": ")
			var response string
			fmt.Scan(&response)

			if strings.Contains(response, "y") {
				count++
				fmt.Printf("Removing %v...\n", file.Name())
				err := os.Remove(file.Name())
				CheckError(err)
			}
		}
	}

	if count == 0 {
		fmt.Println("No tilda files found!")
	}
}

func RemoveFiles(files []fs.FileInfo) {
	count := 0
	for _, file := range files {
		if strings.Contains(file.Name(), "~") {
			count++
			fmt.Printf("Removing %v...\n", file.Name())
			err := os.Remove(file.Name())
			CheckError(err)
		}
	}

	if count == 0 {
		fmt.Println("No tilda files found!")
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
