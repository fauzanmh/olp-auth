package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-auth/entity"
)

type Repository interface {
	// Auth
	CheckUser(ctx context.Context, memberID int64) (bool, error)
	CheckUsername(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, arg *entity.CreateUserParams) (err error)
	DeleteUser(ctx context.Context, arg *entity.DeleteUserParams) error
	GetAdminByUsername(ctx context.Context, username string) (entity.GetAdminByUsernameRow, error)
	GetUserByUsername(ctx context.Context, username string) (entity.GetUserByUsernameRow, error)

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}
