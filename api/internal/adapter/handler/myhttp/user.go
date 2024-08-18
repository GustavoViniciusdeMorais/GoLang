package myhttp

import (
	"net/http"
	"strconv"

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

// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} []domain.User
// @Router /users [get]
func (h *UserHandler) GetUsers(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")
	users, error := h.userService.FindAll(page, limit)

	if error != nil {
		return ctx.JSON(http.StatusInternalServerError, error)
	}

	return ctx.JSON(http.StatusOK, users)
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body UserRequest true "User"
// @Success 200 {object} domain.User
// @Router /users [post]
func (h *UserHandler) CreateUser(ctx echo.Context) error {
	req := new(UserRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	user, err := h.userService.Save(req.Name, req.Email, req.Birthday, string(password))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, user)
}

// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UserRequest true "User"
// @Success 200 {object} domain.User
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	req := new(UserRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	foundUser, err := h.userService.FindById(id)
	if foundUser == nil || err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	if req.Password == "" {
		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password))
		if err != nil {
			newPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, err)
			}
			req.Password = string(newPassword)
		}
	}

	user, err := h.userService.Update(foundUser.ID, req.Name, req.Email, req.Birthday, req.Password, req.Active)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, user)
}
