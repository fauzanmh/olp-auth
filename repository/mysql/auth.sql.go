package mysql

import (
	"context"

	"github.com/fauzanmh/olp-auth/entity"
)

const checkUser = `-- name: CheckUser :one
SELECT EXISTS(SELECT 1 FROM users WHERE member_id = ? LIMIT 1) AS exist
`

func (q *Queries) CheckUser(ctx context.Context, memberID int64) (bool, error) {
	row := q.queryRow(ctx, q.checkUserStmt, checkUser, memberID)
	var exist bool
	err := row.Scan(&exist)
	return exist, err
}

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
INSERT INTO users (username, password, member_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
`

func (q *Queries) CreateUser(ctx context.Context, arg *entity.CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.Password,
		arg.MemberID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users SET deleted_at = ? 
WHERE member_id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, arg *entity.DeleteUserParams) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, arg.DeletedAt, arg.MemberID)
	return err
}
