package main

import (
	"fmt"
	"time"

	"github.com/fdurand/arp"
)

func main() {
	fmt.Println("arp table:")
	for k, v := range arp.Table() {
		fmt.Println(k, "->", v)
	}
	fmt.Printf("last updated at %v", arp.CacheLastUpdate().Format(time.RFC822))
}
