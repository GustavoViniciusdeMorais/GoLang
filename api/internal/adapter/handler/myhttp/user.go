package myhttp

import (
	"net/http"

	"example.com/internal/core/port"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")
	users, error := h.userService.ListUsers(page, limit)

	if error != nil {
		return ctx.JSON(http.StatusInternalServerError, error)
	}

	return ctx.JSON(http.StatusOK, users)
}
