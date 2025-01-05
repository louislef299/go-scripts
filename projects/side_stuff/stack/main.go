package main

import (
	"fmt"
	"log"
	"os"

	"github.com/louislef299/go-scripts/projects/side_stuff/stack/stack"
)

const filename = "stack.log"

func main() {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s, err := stack.NewFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	values := []string{"a", "b", "c", "d"}
	for _, v := range values {
		if err := s.Push(v); err != nil {
			log.Fatal(err)
		}
	}

	val, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
