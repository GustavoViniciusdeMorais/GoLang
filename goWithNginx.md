### Golang API (main.go)

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", helloHandler)
    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}
```

Save this script as `main.go`. To run the server, you'll need to have Go installed. You can then build and run the server using:

```sh
go run main.go
```

The server will start on port 8080.

### NGINX Configuration (nginx.conf)

Below is an example of a simple NGINX configuration file to proxy requests to the Golang API. 

```nginx
server {
    listen 80;

    server_name localhost;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Save this configuration as `nginx.conf`. Make sure to adjust the `server_name` directive to match your domain name or use `localhost` if you're testing locally.

### Running NGINX

1. Install NGINX if you haven't already. You can usually do this using your package manager. For example, on Ubuntu, you can use:

    ```sh
    sudo apt update
    sudo apt install nginx
    ```

2. Replace the default NGINX configuration file with your custom configuration. On many systems, the default configuration file is located at `/etc/nginx/nginx.conf`.

    ```sh
    sudo cp nginx.conf /etc/nginx/nginx.conf
    ```

3. Restart NGINX to apply the new configuration:

    ```sh
    sudo systemctl restart nginx
    ```

### Testing the Setup

With the Go server running and NGINX configured and restarted, you should be able to make a request to the NGINX server on port 80, which will proxy the request to the Go server on port 8080.

Open your browser and navigate to `http://localhost`. You should see "Hello, World!" as the response.

This setup provides a simple reverse proxy configuration using NGINX to forward requests to a Golang API. It’s a basic but effective way to handle requests and can be expanded with additional features as needed.