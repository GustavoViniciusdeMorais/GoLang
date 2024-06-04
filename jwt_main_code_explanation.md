### Imports

```go
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
```

- `github.com/dgrijalva/jwt-go`: Library for creating and parsing JWT tokens.
- `github.com/go-redis/redis/v8`: Library for interacting with a Redis database.
- `github.com/gorilla/mux`: HTTP router and dispatcher for matching incoming requests to their respective handler.
- `net/http`: Standard library for HTTP client and server implementations.
- `time`: Standard library for time manipulation.
- `context`: Standard library for managing context in APIs.
- `fmt`: Standard library for formatted I/O.
- `log`: Standard library for logging.
- `os`: Standard library for interacting with the operating system.

### Global Variables

```go
var jwtKey = []byte("my_secret_key")
var redisClient *redis.Client
var ctx = context.Background()
```

- `jwtKey`: Secret key used to sign JWT tokens.
- `redisClient`: Global variable to hold the Redis client connection.
- `ctx`: Context used for Redis operations.

### Structs

```go
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}
```

- `Credentials`: Represents the username and password for authentication.
- `Claims`: Custom claims for JWT which includes the standard claims and the username.

### main Function

```go
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
```

- Connect to the Redis database running on `redis:6379`.
- Create a new router using Gorilla Mux.
- Define routes for `/signin` (POST) and `/welcome` (GET).
- Start the HTTP server on port 8000.

### Signin Handler

```go
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
```

- Decode the incoming JSON request into the `Credentials` struct.
- Check if the credentials are correct (in a real application, validate against a database).
- Create a JWT token with an expiration time of 5 minutes.
- Sign the JWT token with the secret key.
- Store the JWT token in Redis with the username as the key.
- Set the JWT token in a cookie in the response.

### Welcome Handler

```go
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

- Retrieve the token from the request cookies.
- Parse the token and validate it.
- If the token is invalid or expired, respond with an appropriate error.
- Check if the token exists and matches the token stored in Redis.
- Respond with a welcome message if everything is valid.

This code demonstrates a simple JWT-based authentication system with Golang, using Redis for token storage and Docker for containerization.