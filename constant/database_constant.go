package constant

import "fmt"

type ErrorMysql error

var (
	ErrorMysqlUserNotFound      ErrorMysql = fmt.Errorf("user not found")
	ErrorMysqlUserAlreadyExists ErrorMysql = fmt.Errorf("user already exists")
	ErrorMysqlDataNotFound      ErrorMysql = fmt.Errorf("data not found")
)
