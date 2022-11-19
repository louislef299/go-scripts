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

func Scan(ctx context.Context, hostname string, portrange int) (chan ScanResult, chan struct{}) {
	r := make(chan ScanResult)
	done := make(chan struct{})

	g, ctx := errgroup.WithContext(ctx)
	protocol := "tcp"
	for port := 0; port <= portrange; port++ {
		port := port // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			time.Sleep(time.Second)
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
	fmt.Println("Port Scanning 1024 range on", target)
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()

	long_results, done := Scan(ctx, target, 49152)
	for {
		select {
		case l := <-long_results:
			if isOpen(l.State) {
				fmt.Printf("%s:%d is %s\n", l.Protocol, l.Port, l.State)
			}
		case <-done:
			fmt.Println("scan is finished")
			return
		case <-ctx.Done():
			fmt.Println("recieved SIGINT, exiting")
			return
		}
	}
}
