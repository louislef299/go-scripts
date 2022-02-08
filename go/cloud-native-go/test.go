package main

import (
	"fmt"
	"time"
)

func main() {
	lastAttempt := time.Now()
	for i := 0; i < 10; i++ {
		lastAttempt = lastAttempt.Add(time.Second * 2 << i)
		fmt.Println(time.Second * 2 << i)
		fmt.Println(lastAttempt)
	}
}
