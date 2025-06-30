package aws

import (
	"context"
	"fmt"

	"github.com/louislef299/comptime-login/pkg/login"
)

type AWSLogin struct{}

func init() {
	login.Register("aws", &AWSLogin{})
}

func (a *AWSLogin) Login(ctx context.Context, opts ...login.ConfigOptionsFunc) error {
	fmt.Println("hello from AWS!")
	return nil
}
