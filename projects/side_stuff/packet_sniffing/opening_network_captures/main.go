package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var fname = flag.String("pcap", "", "Pcap File to load and parse")
var filter = flag.String("filter", "", "BPF Filter to apply")

func main() {
	log.Println("start")
	defer log.Println("end")

	flag.Parse()

	// Check if file has been passed, otherwise print help
	if *fname == "" {
		flag.PrintDefaults()
		log.Fatalf("no pcap file parameter was passed")
	}

	pcapHandle, err := pcap.OpenOffline(*fname)
	if err != nil {
		log.Fatalf("error opening file - %v", err)
	}

	defer pcapHandle.Close()

	if *filter != "" {
		err = pcapHandle.SetBPFFilter(*filter)
		if err != nil {
			log.Fatalf("error appling filter %s - %v", *filter, err)
		}
	}

	packetsFiltered := gopacket.NewPacketSource(pcapHandle, pcapHandle.LinkType())

	for packet := range packetsFiltered.Packets() {
		fmt.Println(packet)
	}
}
