package docker

import (
	"context"
	"fmt"

	"github.com/louislef299/comptime-login/pkg/config"
	"github.com/louislef299/comptime-login/pkg/login"
)

type DockerLogin struct{}

func init() {
	login.Register("docker", &DockerLogin{})
}

func (a *DockerLogin) Login(ctx context.Context, config *config.Config, opts ...login.ConfigOptionsFunc) error {
	fmt.Printf("Cluster: %s\n", config.Cluster)
	return nil
}
