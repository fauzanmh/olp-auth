package helper

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/fauzanmh/olp-auth/constant"
	"github.com/lib/pq"
)

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func PqError(err error) (int, error) {
	re := regexp.MustCompile("\\((.*?)\\)")
	if err, ok := err.(*pq.Error); ok {
		match := re.FindStringSubmatch(err.Detail)

		switch err.Code.Name() {
		case "unique_violation":
			return pqErrorMap["unique_violation"], fmt.Errorf("%s already exists", match[1])
		}
	}
	return http.StatusInternalServerError, fmt.Errorf("internal error")
}

var commonErrorMap = map[error]int{
	constant.ErrorMysqlUserAlreadyExists: http.StatusConflict,
	constant.ErrorMysqlUserNotFound:      http.StatusNotFound,
	constant.ErrorMysqlDataNotFound:      http.StatusBadRequest,
	constant.ErrorMessageUsernameUnique:  http.StatusBadRequest,
	constant.ErrorMessageLogin:           http.StatusBadRequest,
}

// CommonError is
func CommonError(err error) (int, error) {
	switch err {
	case constant.ErrorMysqlUserAlreadyExists:
		return commonErrorMap[constant.ErrorMysqlUserAlreadyExists], constant.ErrorMysqlUserAlreadyExists
	case constant.ErrorMysqlUserNotFound:
		return commonErrorMap[constant.ErrorMysqlUserNotFound], constant.ErrorMysqlUserNotFound
	case constant.ErrorMysqlDataNotFound:
		return commonErrorMap[constant.ErrorMysqlDataNotFound], constant.ErrorMysqlDataNotFound
	case constant.ErrorMessageUsernameUnique:
		return commonErrorMap[constant.ErrorMessageUsernameUnique], constant.ErrorMessageUsernameUnique
	case constant.ErrorMessageLogin:
		return commonErrorMap[constant.ErrorMessageLogin], constant.ErrorMessageLogin
	}

	return http.StatusInternalServerError, fmt.Errorf("internal error")
}
