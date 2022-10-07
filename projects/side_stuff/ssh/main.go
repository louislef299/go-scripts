package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

/***
 * Connect over SSH into a running docker container
 *
 * CLI Command:
 * - docker inspect -f "{{ .NetworkSettings.IPAddress }}" container_name
 * - ping first to make sure available:
 *   ping –c 3 172.17.0.2
 * - ssh root@172.17.0.2
 ***/

func main() {
	// ssh config
	hostKeyCallback, err := knownhosts.New("/Users/louis/.ssh/known_hosts")
	if err != nil {
		log.Fatalf("Unable to create host key call back %v", err)
	}

	config := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.Password("test"),
		},
		HostKeyCallback: hostKeyCallback,
	}
	// connect ot ssh server
	conn, err := ssh.Dial("tcp", "127.0.0.1:32000", config)
	if err != nil {
		log.Fatalf("failed to dial the ssh server %v", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// configure terminal mode
	/*modes := ssh.TerminalModes{
		ssh.ECHO: 0, // supress echo
	}
	// run terminal session
	if err := session.RequestPty("xterm", 50, 80, modes); err != nil {
		log.Fatalf("failed to run terminal session %v", err)
	}
	// start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to run remote session %v", err)
	}*/

	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run("ls -la"); err != nil {
		log.Fatalf("failed to send ssh command %v", err)
	}
	fmt.Println(buff.String())
}
