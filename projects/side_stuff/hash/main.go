package main

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)


func main() {
	basic()
	file()
}

func basic() {
	s := "Hello world"

	hsha2 := sha256.Sum256([]byte(s))
	log.Printf("SHA256: %x\n", hsha2)
}

func file() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}

	sum := sha256.Sum(nil)
	log.Printf("File SHA256: %x\n", sum)
}