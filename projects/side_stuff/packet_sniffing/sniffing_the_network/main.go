package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var filter = flag.String("filter", "", "BPF filter for capture")
var iface = flag.String("iface", "en0", "Select interface where to capture")
var snaplen = flag.Int("snaplen", 1024, "Maximun sise to read for each packet")
var promisc = flag.Bool("promisc", false, "Enable promiscuous mode")
var timeoutT = flag.Int("timeout", 30, "Connection Timeout in seconds")
var file = flag.String("file", "", "Write packet output to a file")

func main() {
	log.Println("start")
	defer log.Println("end")

	flag.Parse()

	var timeout time.Duration = time.Duration(*timeoutT) * time.Second

	// Opening Device
	handle, err := pcap.OpenLive(*iface, int32(*snaplen), *promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	// Applying BPF Filter if it exists
	if *filter != "" {
		log.Println("applying filter ", *filter)
		err := handle.SetBPFFilter(*filter)
		if err != nil {
			log.Fatalf("error applyign BPF Filter %s - %v", *filter, err)
		}
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	out := os.Stdout
	if *file != "" {
		f, err := os.OpenFile(*file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		out = f
	}
	for packet := range packetSource.Packets() {
		fmt.Fprintf(out, "%v\n", packet)

		if app := packet.ApplicationLayer(); app != nil {
			fmt.Println("Payload:", string(app.Payload()))
		}
	}
}
