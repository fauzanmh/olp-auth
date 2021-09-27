package auth

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	MemberID int64  `json:"member_id"`
}
