package http

import (
	"github.com/fauzanmh/olp-auth/pkg/util"
	"github.com/fauzanmh/olp-auth/schema/auth"
	usecase "github.com/fauzanmh/olp-auth/usecase/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase usecase.Usecase
}

func NewAuthHandler(e *echo.Group, usecase usecase.Usecase) {
	handler := &AuthHandler{
		usecase: usecase,
	}

	routerV1 := e.Group("/v1")
	routerV1.POST("/user", handler.CreateUser)
}

// CreateUser nodoc
func (h *AuthHandler) CreateUser(c echo.Context) error {
	req := auth.CreateUserRequest{}
	ctx := c.Request().Context()

	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	err = h.usecase.CreateUser(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success create user", nil)
}
