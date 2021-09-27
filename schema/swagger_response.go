package schema

import "github.com/fauzanmh/olp-auth/schema/auth"

type SwaggerLoginResponse struct {
	Base
	Data auth.LoginResponse `json:"data"`
}
