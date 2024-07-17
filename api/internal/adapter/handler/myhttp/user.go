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
	users, error := h.userService.ListUsers()

	if error != nil {
		return ctx.JSON(http.StatusInternalServerError, error)
	}

	return ctx.JSON(http.StatusOK, users)
}
