package myhttp

import (
	"context"
	"net/http"
	"strings"

	"example.com/internal/core/port"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(redisCache port.CacheRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing or invalid token"})
			}

			// Extract the token from the header
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
			}

			// Parse the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userEmail := claims["email"].(string)

				// Get the token from Redis and compare
				ctx := context.Background()
				savedToken, err := redisCache.Get(ctx, userEmail)
				if err != nil || savedToken != tokenString {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
				}

				// Set the user in the context
				c.Set("user", token)
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}
	}
}
