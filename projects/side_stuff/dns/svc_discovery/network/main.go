package main

import (
	"log"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("Oops: %v\n", err)
	}
	log.Printf("hostname: %s", name)

	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Fatalf("Oops: %v\n", err)
	}

	for _, a := range addrs {
		log.Println(a)
	}

	log.Println("Outbound IP:", GetOutboundIP())
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// sudo nmap -sn 192.168.1.0/24
