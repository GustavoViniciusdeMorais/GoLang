package myhttp

import (
	"net/http"
	"time"

	"example.com/internal/core/port"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService port.AuthService
	redisCache  port.CacheRepository
}

func (a *AuthHandler) Login(ctx echo.Context) error {
	user, err := a.authService.Login(ctx.FormValue("email"), ctx.FormValue("password"))
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	userToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}

	err = a.redisCache.Set(ctx.Request().Context(), user.Email, []byte(userToken), 0)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not save token"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": userToken})
}
func (a *AuthHandler) CreateToken(email string) (string, error) {
	return "", nil
}
func (a *AuthHandler) VerifyToken(token string) (string, error) {
	return "", nil
}

func NewAuthHandler(auth port.AuthService, redisCache port.CacheRepository) *AuthHandler {
	return &AuthHandler{
		authService: auth,
		redisCache:  redisCache,
	}
}
