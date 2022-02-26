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

/*** OLD FUNCTION NOT USED ANYMORE REFER TO COBRA DIRECTORY ***/

func main() {
	interactive := flag.Bool("i", false, "Interactive Mode")
	help := flag.Bool("help", false, "Print guide for how to use rmt")
	h := flag.Bool("h", false, "Print guide for how to use rmt")
	flag.Parse()

	files := ListDir()

	if *interactive {
		InterRemoveFiles(files)
	} else if *h || *help {
		Help()
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

func Help() {
	fmt.Println("Removes all files containing a tilda from the current working directory")
	fmt.Printf("\nUsage: rmt [ -i | -help ]\n\n")
	fmt.Println("-i: Interactive mode, type y to delete files")
	fmt.Println("-help | -h: Prints out usage of rmt")
}
