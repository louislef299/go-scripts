package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*** OLD FUNCTION NOT USED ANYMORE REFER TO COBRA DIRECTORY ***/

func main() {
	recover := flag.Bool("recover", false, "What file to recover from boot library")
	r := flag.Bool("r", false, "What file to recover from boot library")
	recoverAll := flag.Bool("recover-all", false, "Recover all files from boot library")
	list := flag.Bool("list", false, "List contents of boot library")
	l := flag.Bool("l", false, "List contents of boot library")
	help := flag.Bool("help", false, "Description of how to use boot")
	h := flag.Bool("h", false, "Description of how to use boot")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: boot [ -recover | -recover-all | -list | -help ] <files>")
		os.Exit(1)
	}

	if *recoverAll {
		BootRecoverAll()
	} else if *recover || *r {
		bootArgs := os.Args[2:]
		BootRecover(bootArgs)
	} else if *list || *l {
		BootList()
	} else if *h || *help {
		Help()
	} else {
		bootArgs := os.Args[1:]
		Boot(bootArgs)
	}
}

func Boot(bootArgs []string) {
	workingDir, err := os.Getwd()
	CheckError(err)

	for _, bootFile := range bootArgs {
		path, err := os.Getwd()
		CheckError(err)
		exists := DoesFileExist(bootFile, path)
		if !exists {
			fmt.Println("Could not find file", bootFile, "in", workingDir, "to boot!")
			continue
		}

		bootPathFrom := fmt.Sprintf("%v/%v", workingDir, bootFile)
		homeDir, err := os.UserHomeDir()
		CheckError(err)
		bootPathTo := fmt.Sprintf("%v/.boot/%v", homeDir, bootFile)
		err = os.Rename(bootPathFrom, bootPathTo)
		CheckError(err)
		fmt.Println("Booted", bootFile, "to", bootPathTo)
	}
}

func BootList() {
	homeDir, err := os.UserHomeDir()
	CheckError(err)
	files, err := ioutil.ReadDir(fmt.Sprintf("%v/.boot", homeDir))
	CheckError(err)

	if len(files) == 0 {
		fmt.Println("There are no files in boot")
		return
	}

	fmt.Println("The following files are in boot:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func BootRecover(bootArgs []string) {
	workingDir, err := os.Getwd()
	CheckError(err)

	for _, bootFile := range bootArgs {
		homeDir, err := os.UserHomeDir()
		CheckError(err)
		bootPath := fmt.Sprintf("%v/.boot", homeDir)
		exists := DoesFileExist(bootFile, bootPath)
		if !exists {
			fmt.Println("Could not find file", bootFile, "in", bootPath, "to recover!")
			continue
		}

		bootPathFrom := fmt.Sprintf("%v/.boot/%v", homeDir, bootFile)
		bootPathTo := fmt.Sprintf("%v/%v", workingDir, bootFile)
		err = os.Rename(bootPathFrom, bootPathTo)
		CheckError(err)
		fmt.Println("Recovered", bootFile)
	}
}

func BootRecoverAll() {
	homeDir, err := os.UserHomeDir()
	CheckError(err)
	files, err := ioutil.ReadDir(fmt.Sprintf("%v/.boot", homeDir))
	CheckError(err)
	workingDir, err := os.Getwd()
	CheckError(err)

	for _, file := range files {
		bootPathFrom := fmt.Sprintf("%v/.boot/%v", homeDir, file.Name())
		bootPathTo := fmt.Sprintf("%v/%v", workingDir, file.Name())
		err := os.Rename(bootPathFrom, bootPathTo)
		CheckError(err)
		fmt.Println("Recovered", file.Name())
	}
}

func DoesFileExist(bootFile string, path string) bool {
	files, err := ioutil.ReadDir(path)
	CheckError(err)

	fileFound := false
	for _, file := range files {
		if file.Name() == bootFile {
			fileFound = true
		}
	}

	return fileFound
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Help() {
	fmt.Printf("The boot command is used to either move a file from the current working directory into the boot directory, or to move a file from the boot directory into the current working directory.\n\n")
	fmt.Printf("Usage: boot [ -recover | -recover-all | -list ] <files>\n\n")
	fmt.Println("Options:")
	fmt.Println("-list | -l: Lists the contents of the boot directory")
	fmt.Println("-recover | -r: Recover files boot directory into current working directory")
	fmt.Println("-recover-all: Recover all files from boot into current working directory")
	fmt.Println("-help | -h: List out the command options")
}
