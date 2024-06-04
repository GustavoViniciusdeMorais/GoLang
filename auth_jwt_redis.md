## JWT with Redis

### 1. Directory Structure

```
go-redis-jwt/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

### 2. main.go
- [Code details](./jwt_main_code_explanation.md)
```go
package main

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/go-redis/redis/v8"
    "github.com/gorilla/mux"
    "net/http"
    "time"
    "context"
    "fmt"
    "log"
    "os"
)

var jwtKey = []byte("my_secret_key")
var redisClient *redis.Client
var ctx = context.Background()

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func main() {
    // Connect to Redis
    redisClient = redis.NewClient(&redis.Options{
        Addr: "redis:6379",
        Password: "",
        DB: 0,
    })

    router := mux.NewRouter()

    router.HandleFunc("/signin", Signin).Methods("POST")
    router.HandleFunc("/welcome", Welcome).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}

func Signin(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // This is a mock. You should check user credentials from a DB.
    if creds.Username != "user" || creds.Password != "password" {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        Username: creds.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Store JWT in Redis
    err = redisClient.Set(ctx, creds.Username, tokenString, expirationTime.Sub(time.Now())).Err()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })
}

func Welcome(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    tokenStr := cookie.Value
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Validate token with Redis
    storedToken, err := redisClient.Get(ctx, claims.Username).Result()
    if err != nil || storedToken != tokenStr {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
```

### 3. Dockerfile

```dockerfile
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-redis-jwt

EXPOSE 8000

CMD ["/go-redis-jwt"]
```

### 4. docker-compose.yml

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8000:8000"
    depends_on:
      - redis
  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
```

### 5. go.mod

```go
module go-redis-jwt

go 1.18

require (
    github.com/dgrijalva/jwt-go v3.2.0+incompatible
    github.com/go-redis/redis/v8 v8.11.4
    github.com/gorilla/mux v1.8.0
)
```

### Steps to Run

1. **Navigate to your project directory:**

    ```sh
    cd go-redis-jwt
    ```

2. **Build and run the Docker containers using Docker Compose:**

    ```sh
    docker-compose up --build
    ```

3. **Test the API:**

    - Sign in to get a token:

        ```sh
        curl -X POST -d '{"username":"user", "password":"password"}' -H "Content-Type: application/json" http://localhost:8000/signin -c cookies.txt
        ```

    - Access the protected endpoint:

        ```sh
        curl -b cookies.txt http://localhost:8000/welcome
        ```
