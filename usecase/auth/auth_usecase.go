package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/fauzanmh/olp-auth/constant"
	"github.com/fauzanmh/olp-auth/entity"
	appInit "github.com/fauzanmh/olp-auth/init"
	"github.com/fauzanmh/olp-auth/pkg/util"
	mysqlRepo "github.com/fauzanmh/olp-auth/repository/mysql"
	"github.com/fauzanmh/olp-auth/schema/auth"
)

type usecase struct {
	config    *appInit.Config
	mysqlRepo mysqlRepo.Repository
}

func NewAuthUseCase(config *appInit.Config, mysqlRepo mysqlRepo.Repository) Usecase {
	return &usecase{
		config:    config,
		mysqlRepo: mysqlRepo,
	}
}

// --- create user --- //
func (u *usecase) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (err error) {
	// check username is exists
	exist, err := u.mysqlRepo.CheckUsername(ctx, req.Username)
	if err != nil {
		return
	}
	if exist {
		err = constant.ErrorMessageUsernameUnique
		return
	}

	// arguments
	password, err := util.HashPassword(req.Password)
	if err != nil {
		return
	}
	createUserParams := &entity.CreateUserParams{
		Username:  req.Username,
		Password:  password,
		MemberID:  req.MemberID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}

	// create user to database
	err = u.mysqlRepo.CreateUser(ctx, createUserParams)
	if err != nil {
		return
	}

	return
}

// --- delete user --- //
func (u *usecase) DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) (err error) {
	// check user is exists
	exist, err := u.mysqlRepo.CheckUser(ctx, req.MemberID)
	if err != nil {
		return
	}
	if !exist {
		err = constant.ErrorMessageUserNotFound
		return
	}

	// delete user from database
	deleteUserParams := &entity.DeleteUserParams{
		MemberID:  req.MemberID,
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	}
	err = u.mysqlRepo.DeleteUser(ctx, deleteUserParams)
	if err != nil {
		return
	}

	return
}
