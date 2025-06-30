package aws

import (
	"context"

	"github.com/louislef299/comptime-login/pkg/config"
	"github.com/louislef299/comptime-login/pkg/login"
)

type AWSLogin struct{}

func init() {
	login.Register("aws", &AWSLogin{})
}

func (a *AWSLogin) Login(ctx context.Context, config *config.Config, opts ...login.ConfigOptionsFunc) error {
	config.Role = "aws:arn:louis"
	config.Secret = "poopybutt"
	return nil
}
