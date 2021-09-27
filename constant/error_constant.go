package constant

import "fmt"

type ErrorMessage error

var (
	ErrorMessageUserNotFound   ErrorMessage = fmt.Errorf("user not found")
	ErrorMessageUsernameUnique ErrorMessage = fmt.Errorf("username has been taken")
	ErrorMessageLogin          ErrorMessage = fmt.Errorf("username / password is wrong")
)
