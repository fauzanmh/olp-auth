package entity

import "database/sql"

// model
type User struct {
	ID        int64         `json:"id"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
	DeletedAt sql.NullInt64 `json:"deleted_at"`
}

// params and rows //
type CreateUserParams struct {
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	MemberID  int64         `json:"member_id"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
}

type DeleteUserParams struct {
	DeletedAt sql.NullInt64 `json:"deleted_at"`
	MemberID  int64         `json:"member_id"`
}

type GetUserByUsernameRow struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
