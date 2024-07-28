package myhttp

import (
	"net/http"

	"example.com/internal/core/port"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	req := new(CreateUserRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	user, err := h.userService.CreateUser(req.Name, req.Email, req.Birthday, string(password))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, user)
}
