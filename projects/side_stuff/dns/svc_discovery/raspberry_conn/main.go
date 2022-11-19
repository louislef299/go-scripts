package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"strings"

	"github.com/jpillora/icmpscan"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

var (
	verbose bool
	piname  string
	homedir string
	sshkey  string
)

func interactiveSession(session *ssh.Session) {
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

func printhost(h *icmpscan.Host) {
	if h.Hostname != "" {
		fmt.Printf("%v has hostname %s\n", h.IP, h.Hostname)
	} else {
		fmt.Printf("%v does not have a hostname\n", h.IP)
	}
}

func init() {
	var err error
	homedir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal("could not get home directory:", err)
	}

	flag.BoolVar(&verbose, "v", false, "verbose output of all connections")
	flag.StringVar(&piname, "piname", "raspberrypi", "the dns name of the pi")
	flag.StringVar(&sshkey, "sshkey", ".ssh/id_rsa", "the path to the ssh public key to use")
}

func main() {
	flag.Parse()
	log.Println("gathering all hosts from the local network")
	hosts, err := icmpscan.Run(icmpscan.Spec{
		Hostnames: true,
		MACs:      true,
		Log:       verbose,
	})
	if err != nil {
		log.Fatal("could not scan:", err)
	}

	log.Println("locating pi ip address")
	var ip net.IP
	found := false
	for _, host := range hosts {
		if verbose {
			printhost(host)
		}
		if strings.Compare(host.Hostname, piname) == 0 {
			ip = host.IP
			found = true
			break
		}
	}
	if !found {
		log.Fatal("could not find pi! add verbose flag to see list of hosts")
	}

	log.Println("generating ssh configuration")
	hostKeyCallback, err := knownhosts.New(path.Join(homedir, ".ssh/known_hosts"))
	if err != nil {
		log.Fatalf("Unable to create host key call back %v", err)
	}

	key, err := os.ReadFile(path.Join(homedir, sshkey))
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	config := &ssh.ClientConfig{
		User: "louis",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	target := ip.String() + ":22"
	client, err := ssh.Dial("tcp", target, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	log.Println("starting ssh session")
	interactiveSession(session)
}
