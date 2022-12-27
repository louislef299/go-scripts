package stack

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	file *os.File
}

func NewFile(filename string) (*Stack, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &Stack{
		file: f,
	}, nil
}

func (s *Stack) Push(value string) error {
	_, err := fmt.Fprint(s.file, value)
	return fmt.Errorf("could not push: %v", err)
}

func (s *Stack) Pop() (string, error) {
	r := bufio.NewReader(s.file)
	fmt.Println(r.Size())
	return "", fmt.Errorf("there was an issue")
}
