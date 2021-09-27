package mysql

import (
	"context"
	"database/sql"

	"github.com/fauzanmh/olp-auth/entity"
)

type Repository interface {
	// Auth
	CheckUsername(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, arg *entity.CreateUserParams) (err error)

	//Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	WithTx(tx *sql.Tx) *Queries
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}
