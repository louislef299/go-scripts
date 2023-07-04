package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("error retrieving devices - %v", err)
	}

	for _, device := range devices {
		fmt.Printf("Device Name: %s\n", device.Name)
		fmt.Printf("Device Description: %s\n", device.Description)
		fmt.Printf("Device Flags: %d\n", device.Flags)
		for _, iaddress := range device.Addresses {
			fmt.Printf("\tInterface IP: %s\n", iaddress.IP)
			fmt.Printf("\tInterface NetMask: %s\n", iaddress.Netmask)
		}
	}

	// A more modern way of looking for network interfaces
	d, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range d {
		fmt.Println(device.Name)
	}
}
