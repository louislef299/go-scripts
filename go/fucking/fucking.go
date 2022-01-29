package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	CheckError(err)

	imageName := "docker.io/library/ubuntu"

	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	CheckError(err)

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:       "ubuntu",
		Tty:         true,
		AttachStdin: true,
	}, nil, nil, nil, "")
	CheckError(err)

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	CheckError(err)

	command := []string{"/bin/bash"}

	configCreate := types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          command,
	}

	respC, err := cli.ContainerExecCreate(ctx, resp.ID, configCreate)
	CheckError(err)
	fmt.Println("Creation response:", respC.ID)

	config := types.ExecStartCheck{
		Detach: true,
		Tty:    true,
	}

	time.Sleep(2)

	respA, err := cli.ContainerExecAttach(ctx, respC.ID, config)
	CheckError(err)

	fmt.Println("Response output:", respA)

	/*
		// Take this out to run in the background
		statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			CheckError(err)
		case <-statusCh:
		}

		out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
		CheckError(err)

		stdcopy.StdCopy(os.Stdout, os.Stderr, out)*/
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
