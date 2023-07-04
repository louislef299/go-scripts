package main

import (
	"fmt"
	"log"

	"github.com/jpillora/icmpscan"
)

// sudo go run main.go
func main() {
	hosts, err := icmpscan.Run(icmpscan.Spec{
		Hostnames: true,
		MACs:      true,
		Log:       false,
	})
	if err != nil {
		log.Fatal("could not scan:", err)
	}
	for _, host := range hosts {
		fmt.Printf("%v has hostname %s\n", host.IP, host.Hostname)
	}
}
