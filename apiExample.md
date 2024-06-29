# Setting Up a Simple Go API Project with a Single `go.mod` File

This tutorial will guide you through setting up a simple Go API project with the following directory structure:

```
./
├── go.mod
├── internal
│   └── server
│       └── server.go
└── main.go
```

## Step 1: Create the Project Directory Structure

First, create the project directory structure. You can use the following commands to set it up:

```sh
mkdir -p myproject/internal/server
cd myproject
touch go.mod main.go internal/server/server.go
```

Your directory structure should now look like this:

```
myproject/
├── go.mod
├── internal
│   └── server
│       └── server.go
└── main.go
```

## Step 2: Initialize the Go Module

In the root directory of your project (`myproject`), initialize the Go module by running:

```sh
go mod init example.com/myproject
```

This command will create the `go.mod` file with the module path `example.com/myproject`.

## Step 3: Write the `main.go` File

Open the `main.go` file and add the following code:

```go
package main

import (
    "fmt"
    "net/http"
    "example.com/myproject/internal/server"
)

func main() {
    http.HandleFunc("/", server.HandleRoot)
    fmt.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
```

This code sets up a basic HTTP server that uses a handler from the `internal/server` package.

## Step 4: Write the `server.go` File

Open the `internal/server/server.go` file and add the following code:

```go
package server

import (
    "fmt"
    "net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

This code defines a simple HTTP handler that responds with "Hello, World!" when the root URL is accessed.

## Step 5: Tidy Up Dependencies

Run the following command to tidy up the dependencies and ensure the `go.mod` file is up-to-date:

```sh
go mod tidy
```

## Step 6: Run the Server

Start the server by running:

```sh
go run main.go
```

You should see the message `Server is running on port 8080` in your terminal. Open your web browser and go to `http://localhost:8080`. You should see `Hello, World!` displayed.

## Summary

You have successfully set up a simple Go API project with a single `go.mod` file. The project includes a basic HTTP server that handles requests to the root URL. This setup provides a solid foundation for building more complex Go applications.

## Complete Directory Structure

After completing the tutorial, your directory structure should look like this:

```
myproject/
├── go.mod
├── internal
│   └── server
│       └── server.go
└── main.go
```

Feel free to expand on this foundation by adding more features and functionality to your Go API project!
