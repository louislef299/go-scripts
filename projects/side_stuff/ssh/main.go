package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
 *
 * Resources:
 * https://linuxhint.com/golang-ssh-examples/
 * https://github.com/inatus/ssh-client-go
 * https://docs.docker.com/engine/reference/commandline/run/#publish-or-expose-port--p---expose
 ***/

func main() {
	fmt.Print("Remote host? (Default=localhost): ")
	server := scanConfig()
	if server == "" {
		server = "127.0.0.1"
	}
	fmt.Print("Port? (Default=22): ")
	port := scanConfig()
	if port == "" {
		port = "22"
	}
	server = server + ":" + port
	fmt.Print("UserName?: ")
	user := scanConfig()

	hostKeyCallback, err := knownhosts.New("/Users/louis/.ssh/known_hosts")
	if err != nil {
		log.Fatalf("Unable to create host key call back %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password("test"),
		},
		HostKeyCallback: hostKeyCallback,
	}
	conn, err := ssh.Dial("tcp", server, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	defer conn.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := conn.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Set IO
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	in, _ := session.StdinPipe()

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}

	// Accepting commands
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)
	}

}

func scanConfig() string {
	config, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	config = strings.Trim(config, "\n")
	return config
}

/*func main() {
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
	}

	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run("ls -la"); err != nil {
		log.Fatalf("failed to send ssh command %v", err)
	}
	fmt.Println(buff.String())
}*/
