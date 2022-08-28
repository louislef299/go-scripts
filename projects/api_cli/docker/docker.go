package docker

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/docker/docker/client"
)

// `open --background -a Docker` to start docker from command line
// death by ports shell `lsof -t -i tcp:80 | xargs kill`
// rm all docker containers `docker ps -a --format "{{.ID}}" | xargs docker rm`

func Init() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	ping, err := cli.Ping(ctx)
	if err != nil {
		log.Printf("failed on os %v:\n%v", runtime.GOOS, err)
		startup(time.Minute)
	}
	fmt.Println(ping)
}

func startup(timeout time.Duration) {
	var cmd *exec.Cmd
	os := runtime.GOOS
	switch os {
	case "darwin":
		cmd = exec.Command("open", "--background", "-a", "Docker")
	case "linux":
		cmd = nil
	}

	d := make(chan error)
	// timeout pattern
	end := make(chan struct{})
	go func() {
		done := time.Now().Add(timeout)
		for {
			if time.Now().After(done) {
				close(end)
				return
			} else {
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		fmt.Println("starting up docker...")
		err := cmd.Run()
		if err != nil {
			d <- err
			return
		}
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		for {
			_, err := cli.Ping(ctx)
			if err != nil {
				time.Sleep(time.Second)
				continue
			} else {
				close(d)
			}
		}
	}()

	select {
	case <-end:
		log.Fatal("timeout occurred waiting for docker desktop")
	case r := <-d:
		if r != nil {
			log.Fatalf("error pinging docker desktop: %v", r)
		}
		return
	}
}
