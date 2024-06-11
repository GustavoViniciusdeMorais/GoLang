# Hexagonal API Example

### Project Structure
```
myapp/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── adapters/
│   │   └── http/
│   │       ├── handler.go
│   │       └── router.go
│   ├── application/
│   │   └── service/
│   │       └── user_service.go
│   └── domain/
│       └── user.go
└── go.mod
```

### Step-by-Step Implementation

#### 1. Create the `User` Entity
**internal/domain/user.go**
```go
package domain

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

#### 2. Create the User Service
**internal/application/service/user_service.go**
```go
package service

import "myapp/internal/domain"

type UserService struct{}

func NewUserService() *UserService {
    return &UserService{}
}

func (s *UserService) GetUsers() []domain.User {
    // Mockup data
    return []domain.User{
        {ID: 1, Name: "John Doe", Age: 30},
        {ID: 2, Name: "Jane Doe", Age: 25},
    }
}
```

#### 3. Create the HTTP Handler
**internal/adapters/http/handler.go**
```go
package http

import (
    "encoding/json"
    "net/http"

    "myapp/internal/application/service"
)

type Handler struct {
    userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
    return &Handler{userService: userService}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users := h.userService.GetUsers()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
```

#### 4. Create the Router
**internal/adapters/http/router.go**
```go
package http

import (
    "net/http"

    "github.com/gorilla/mux"
)

func NewRouter(handler *Handler) *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/users", handler.GetUsers).Methods("GET")
    return router
}
```

#### 5. Setup the Main Server File
**cmd/server/main.go**
```go
package main

import (
    "log"
    "net/http"

    "myapp/internal/adapters/http"
    "myapp/internal/application/service"
)

func main() {
    userService := service.NewUserService()
    handler := http.NewHandler(userService)
    router := http.NewRouter(handler)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
```

#### 6. Initialize the Go Module
Run the following command in the root directory of your project to initialize a new Go module:
```sh
go mod init myapp
go get -u github.com/gorilla/mux
```

### Summary
This setup ensures a clear separation of concerns:
- `domain` package contains the core entities.
- `application/service` package contains the business logic.
- `adapters/http` package contains the HTTP handler and router.
- `cmd/server` package contains the main entry point for the application.

To run the server, simply execute the following command in the root directory:
```sh
go run cmd/server/main.go
```
This will start the server on port `8080`, and you can access the users endpoint at `http://localhost:8080/users`.