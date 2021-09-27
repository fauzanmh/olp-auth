package auth

import (
	"context"

	"github.com/fauzanmh/olp-auth/schema/auth"
)

type Usecase interface {
	CreateUser(ctx context.Context, req *auth.CreateUserRequest) (err error)
	DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) (err error)
	Login(ctx context.Context, req *auth.LoginRequest) (res auth.LoginResponse, err error)
}
