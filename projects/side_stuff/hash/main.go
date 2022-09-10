package main

import (
	"crypto/sha256"
	"fmt"
)


func main() {
	s := "Hello world"

	hsha2 := sha256.Sum256([]byte(s))
	fmt.Printf("SHA256: %x\n", hsha2)
}