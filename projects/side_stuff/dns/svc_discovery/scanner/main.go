package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	closed = "Closed"
	open   = "Open"
)

type ScanResult struct {
	Port     int
	State    string
	Protocol string
}

func isOpen(state string) bool {
	if strings.Compare(state, open) == 0 {
		return true
	}
	return false
}

func ScanPort(protocol, hostname string, port int, r chan ScanResult, wg *sync.WaitGroup) {
	defer wg.Done()
	result := ScanResult{Port: port, Protocol: protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		result.State = closed
		r <- result
		return
	}
	defer conn.Close()
	result.State = open
	r <- result
}

func Scan(hostname string, portrange int) chan ScanResult {
	results := make(chan ScanResult)

	var wg sync.WaitGroup
	for i := 0; i <= portrange; i++ {
		wg.Add(2)
		go ScanPort("udp", hostname, i, results, &wg)
		go ScanPort("tcp", hostname, i, results, &wg)
	}

	go func(c chan ScanResult) {
		wg.Wait()
		close(results)
	}(results)

	return results
}

func main() {
	target := "localhost"
	fmt.Println("Port Scanning 1024 range on", target)
	results := Scan(target, 1024)
	for s := range results {
		if isOpen(s.State) {
			fmt.Printf("%s:%d is %s\n", s.Protocol, s.Port, s.State)
		}
	}

	fmt.Println("Port Scanning 49152 range on", target)
	results = Scan(target, 49152)
	for s := range results {
		if isOpen(s.State) {
			fmt.Printf("%s:%d is %s\n", s.Protocol, s.Port, s.State)
		}
	}
}
