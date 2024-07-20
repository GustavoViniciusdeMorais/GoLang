# Golang Auth with Echo Jwt and Redis

### 1. `main.go`
This file will initialize the Echo server, connect to the database and Redis, and set up the routes.

```go
package main

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"

	"your_project/auth"
	"your_project/routes"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Set up routes
	routes.SetupRoutes(e, db, rdb)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
```

### 2. `routes.go`
This file will define the routes for the application.

```go
package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"your_project/auth"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, rdb *redis.Client) {
	authHandler := auth.NewHandler(db, rdb)

	e.POST("/login", authHandler.Login)

	r := e.Group("/welcome")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", authHandler.Welcome)
}
```

### 3. `auth_handler.go`
This file will handle the authentication routes.

```go
package auth

import (
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type Handler struct {
	DB  *gorm.DB
	RDB *redis.Client
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewHandler(db *gorm.DB, rdb *redis.Client) *Handler {
	return &Handler{
		DB:  db,
		RDB: rdb,
	}
}

func (h *Handler) Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var user User
	if err := h.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}

	// Save JWT token in Redis
	ctx := context.Background()
	err = h.RDB.Set(ctx, user.Username, t, 0).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not save token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}

func (h *Handler) Welcome(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome " + name})
}
```

### 4. `auth_service.go`
This file is not strictly necessary for this simple example, but it could be used to add additional services or business logic related to authentication.

```go
package auth

// Add any additional authentication services here
```

### Directory Structure

```
your_project/
|-- main.go
|-- routes.go
|-- auth/
    |-- auth_handler.go
    |-- auth_service.go
```

### Notes
1. Ensure you have the necessary dependencies in your `go.mod` file:

```go
require (
    github.com/go-redis/redis/v8 v8.11.5
    github.com/golang-jwt/jwt v3.2.1+incompatible
    github.com/jinzhu/gorm v1.9.16
    github.com/labstack/echo/v4 v4.6.0
    gorm.io/driver/sqlite v1.2.3
    gorm.io/gorm v1.21.15
)
```

2. This example uses SQLite for simplicity. You can replace it with any other database supported by GORM.

3. Remember to hash passwords in a real application to avoid storing plain text passwords.

This setup provides a basic structure for handling JWT authentication with Golang using the Echo framework, Redis, and GORM.
