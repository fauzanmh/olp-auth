package auth

import (
	"context"

	"github.com/fauzanmh/olp-auth/schema/auth"
)

type Usecase interface {
	CreateUser(ctx context.Context, req *auth.CreateUserRequest) (err error)
}
