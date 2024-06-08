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

### NGINX Configuration

1. Edit the NGINX configuration file:

```sh
sudo nano /etc/nginx/sites-available/api.gustavo.services
```

2. Add the following content to the file:

```nginx
server {
    listen 80;

    server_name api.gustavo.services;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

3. Create a symbolic link to enable the site:

```sh
sudo ln -s /etc/nginx/sites-available/api.gustavo.services /etc/nginx/sites-enabled/
```

4. Test the NGINX configuration for syntax errors:

```sh
sudo nginx -t
```

5. Restart NGINX to apply the new configuration:

```sh
sudo systemctl restart nginx
```

### DNS Configuration

Make sure that your domain `api.gustavo.services` is pointed to the IP address of your NGINX server. This step is usually done through your domain registrar's DNS management interface.
### Testing the Setup

With the Go server running and NGINX configured and restarted, you should be able to make a request to the NGINX server on port 80, which will proxy the request to the Go server on port 8080.

Open your browser and navigate to `http://localhost`. You should see "Hello, World!" as the response.

This setup provides a simple reverse proxy configuration using NGINX to forward requests to a Golang API. It’s a basic but effective way to handle requests and can be expanded with additional features as needed.
