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
	routerV1.POST("/login", handler.Login)

	routerV1.POST("/user", handler.CreateUser)
	routerV1.DELETE("/user/:member_id", handler.DeleteUser)
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

// DeleteUser nodoc
func (h *AuthHandler) DeleteUser(c echo.Context) error {
	req := auth.DeleteUserRequest{}
	ctx := c.Request().Context()

	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	err = h.usecase.DeleteUser(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success delete user", nil)
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body auth.LoginRequest{} true "Request Body"
// @Success 200 {object} schema.SwaggerLoginResponse
// @Failure 400 {object} schema.Base
// @Failure 401 {object} schema.Base
// @Failure 404 {object} schema.Base
// @Failure 500 {object} schema.Base
// @Router /v1/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	req := auth.LoginRequest{}
	ctx := c.Request().Context()

	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	data, err := h.usecase.Login(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success login", data)
}
