package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/sethvargo/go-signalcontext"
	"golang.org/x/sync/errgroup"
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

func ScanWithoutRoutines(hostname string, portrange int) {
	protocol := "tcp"
	for port := 0; port <= portrange; port++ {
		result := ScanResult{Port: port, Protocol: protocol}
		address := hostname + ":" + strconv.Itoa(port)
		conn, err := net.DialTimeout(protocol, address, 60*time.Second)
		if err == nil {
			fmt.Printf("%s:%d is %s\n", result.Protocol, result.Port, result.State)
			conn.Close()
		}
	}
}

func Scan(ctx context.Context, hostname string, portrange int) (chan ScanResult, chan struct{}) {
	r := make(chan ScanResult)
	done := make(chan struct{})

	g, ctx := errgroup.WithContext(ctx)
	protocol := "tcp"
	for port := 0; port <= portrange; port++ {
		port := port // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			result := ScanResult{Port: port, Protocol: protocol}
			address := hostname + ":" + strconv.Itoa(port)
			conn, err := net.DialTimeout(protocol, address, 60*time.Second)
			if err != nil {
				result.State = closed
				r <- result
				return err
			}
			defer conn.Close()
			result.State = open
			r <- result
			return nil
		})
	}

	go func() {
		g.Wait()
		close(r)
		done <- struct{}{}
	}()

	return r, done
}

func main() {
	target := "localhost"
	fmt.Println("Port Scanning 49152 range on", target)
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()

	start := time.Now()
	ScanWithoutRoutines(target, 49152)
	fmt.Printf("scan without goroutines took %v\n", time.Since(start).Truncate(time.Second))
	time.Sleep(time.Second * 3)

	start = time.Now()
	results, done := Scan(ctx, target, 49152)
	for {
		select {
		case r := <-results:
			if isOpen(r.State) {
				fmt.Printf("%s:%d is Open\n", r.Protocol, r.Port)
			}
		case <-done:
			fmt.Printf("scan with goroutines took %v\n", time.Since(start).Truncate(time.Second))
			return
		case <-ctx.Done():
			fmt.Println("recieved SIGINT, exiting")
			return
		}
	}
}
