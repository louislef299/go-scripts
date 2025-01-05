package main

import (
	"fmt"
	"os"
	"time"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	obj := conn.Object("org.example.HelloWorld", "/org/example/HelloWorld")

	// Sending a signal with the interface "org.example.HelloWorld" and method "SayHello"
	err = obj.Emit("org.example.HelloWorld.SayHello", "Hello, World!")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to emit signal:", err)
		os.Exit(1)
	}

	// Let the signal be emitted for a while before exiting
	time.Sleep(1 * time.Second)
}
