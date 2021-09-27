package mysql

import (
	"context"

	"github.com/fauzanmh/olp-auth/entity"
)

const checkUsername = `-- name: CheckUsername :one
SELECT EXISTS(SELECT 1 FROM users WHERE username = ? LIMIT 1) AS exist
`

func (q *Queries) CheckUsername(ctx context.Context, username string) (bool, error) {
	row := q.queryRow(ctx, q.checkUsernameStmt, checkUsername, username)
	var exist bool
	err := row.Scan(&exist)
	return exist, err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (username, password, created_at, updated_at)
VALUES (?, ?, ?, ?)
`

func (q *Queries) CreateUser(ctx context.Context, arg *entity.CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}
