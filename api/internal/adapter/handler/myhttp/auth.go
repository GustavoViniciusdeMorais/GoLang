package myhttp

import (
	"net/http"
	"time"

	"example.com/internal/core/port"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	authService port.AuthService
	redisCache  port.CacheRepository
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	Email string `json:"email"`
}

func (a *AuthHandler) Login(ctx echo.Context) error {
	var loginRequest LoginRequest
	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	user, err := a.authService.FindByEmail(loginRequest.Email)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid password"})
	}

	userToken, err := a.CreateToken(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}

	err = a.redisCache.Set(ctx.Request().Context(), user.Email, []byte(userToken), 0)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not save token"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": userToken})
}

func (a *AuthHandler) Logout(ctx echo.Context) error {
	var logoutRequest LogoutRequest
	if err := ctx.Bind(&logoutRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	_, err := a.authService.FindByEmail(logoutRequest.Email)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "User not found"})
	}

	// Remove the token from Redis
	err = a.redisCache.Delete(ctx.Request().Context(), logoutRequest.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not remove token"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Logout successful"})
}

func (a *AuthHandler) CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString([]byte("secret"))
}

func NewAuthHandler(auth port.AuthService, redisCache port.CacheRepository) *AuthHandler {
	return &AuthHandler{
		authService: auth,
		redisCache:  redisCache,
	}
}
