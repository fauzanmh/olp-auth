package constant

import "fmt"

type ErrorMessage error

var (
	ErrorMessageUsernameUnique ErrorMessage = fmt.Errorf("username has been taken")
)
