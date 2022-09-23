package main

import (
	"flag"
	"log"
	"net"
)

var lookup_url string

func main() {
	flag.StringVar(&lookup_url, "url", "google.com", "The url to lookup")
	flag.Parse()

	ips, err := net.LookupIP(lookup_url)
	if err != nil {
		log.Fatalf("Could not get IPs: %v", err)
	}
	for _, ip := range ips {
		log.Printf("%s. IN A %s\n", lookup_url, ip.String())
	}
}
