# Swagger Documentation with Echo Framework

This guide covers how to set up Swagger documentation for an API built using the Echo framework in Go.

## 1. Setup Swagger in Your Go Project

### Step 1: Install `swag` CLI

Install the `swag` command-line tool to generate Swagger documentation:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Ensure that your `$GOPATH/bin` or `$GOBIN` is in your system's `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Step 2: Add Swagger Annotations

Add Swagger annotations to your handler functions in the `router.go` file:

```go
// @Summary Login a user
// @Description Authenticates a user with their credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param request body domain.LoginRequest true "Login credentials"
// @Success 200 {object} domain.AuthResponse
// @Failure 401 {object} domain.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
    // Handler logic here
}

// @Summary Get all users
// @Description Returns a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User
// @Failure 401 {object} domain.ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetUsers(c echo.Context) error {
    // Handler logic here
}
```

### Step 3: Generate Swagger Documentation

Run the `swag init` command, pointing to your `router.go` file where the routes are defined:

```bash
swag init -g api/internal/adapter/handler/myhttp/router.go
```

This will generate a `docs` folder containing the Swagger documentation.

### Step 4: Import and Serve Swagger Documentation

In your `router.go`, import the generated `docs` package and set up the Swagger route:

```go
import (
    _ "example.com/api/docs" // Replace with your module path
    echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *EchoServer) RegisterRoutes(
    redisCache port.CacheRepository,
    userHandler *UserHandler,
    authHandler *AuthHandler,
) error {
    s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

    ag := s.echo.Group("/auth")
    ag.POST("/login", authHandler.Login)
    ag.POST("/logout", authHandler.Logout)

    ug := s.echo.Group("/users")
    ug.Use(JWTMiddleware(redisCache))
    ug.GET("", userHandler.GetUsers)
    ug.POST("", userHandler.CreateUser)
    ug.PUT("/:id", userHandler.UpdateUser)

    return nil
}
```

### Step 5: View the Documentation

Run your application and visit `http://localhost:8080/swagger/index.html` to view the Swagger UI with your documented API routes.

## 2. Summary of Commands

- Install `swag` CLI:
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```
- Generate Swagger documentation:
  ```bash
  swag init -g api/internal/adapter/handler/myhttp/router.go
  ```

## 3. Swagger Annotations Example

Here’s an example of Swagger annotations for an endpoint:

```go
// @Summary Create a new user
// @Description Adds a new user to the system
// @Tags users
// @Accept json
// @Produce json
// @Param request body domain.CreateUserRequest true "New user details"
// @Success 201 {object} domain.User
// @Failure 400 {object} domain.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
    // Handler logic here
}
```
