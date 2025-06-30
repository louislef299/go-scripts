package k8s

import (
	"context"
	"fmt"

	"github.com/louislef299/comptime-login/pkg/config"
	"github.com/louislef299/comptime-login/pkg/login"
)

type K8sLogin struct{}

func init() {
	login.Register("k8s", &K8sLogin{})
}

func (a *K8sLogin) Login(ctx context.Context, config *config.Config, opts ...login.ConfigOptionsFunc) error {
	fmt.Printf("Role: %s\tSecret: %s\n", config.Role, config.Secret)

	config.Cluster = "louis-dev"
	return nil
}
