package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	if err = conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/org/example/HelloWorld"),
		dbus.WithMatchInterface("org.example.HelloWorld"),
	); err != nil {
		panic(err)
	}

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	// Print each signal received
	for v := range c {
		// Assuming the signal has only one argument: a string
		if len(v.Body) > 0 {
			if msg, ok := v.Body[0].(string); ok {
				fmt.Println("Received message:", msg)
			}
		}
	}
}
