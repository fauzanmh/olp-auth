package auth

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	MemberID int64  `json:"member_id"`
}

type DeleteUserRequest struct {
	MemberID int64 `param:"member_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Provider string `json:"provider"`
}
