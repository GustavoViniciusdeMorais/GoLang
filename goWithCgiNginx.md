### Step-by-Step Configuration

#### 1. Install `fcgiwrap`

`fcgiwrap` is a simple FastCGI wrapper for CGI scripts. You can install it using your package manager. For example, on Ubuntu:

```sh
sudo apt update
sudo apt install fcgiwrap
```

Ensure `fcgiwrap` is running. On many systems, it can be started as a service:

```sh
sudo systemctl start fcgiwrap
sudo systemctl enable fcgiwrap
```

#### 2. Write the Golang Script

Here's an example of a simple Go script that uses the CGI package to work with FastCGI:

```go
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", helloHandler)
    http.Serve(nil, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func init() {
    http.HandleFunc("/", helloHandler)
    err := http.ListenAndServe("127.0.0.1:9000", nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
```

Build the Go script:

```sh
go build -o hello.cgi
```

#### 3. Configure NGINX

Here's an example `nginx.conf` configuration to run the Go script via FastCGI:

```nginx
server {
    listen 80;

    server_name localhost;

    location / {
        fastcgi_pass 127.0.0.1:9000;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME /path/to/your/go/script/hello.cgi;
    }
}
```

Replace `/path/to/your/go/script/hello.cgi` with the actual path to the compiled Go script.

#### 4. Permissions

Ensure that the `hello.cgi` script is executable:

```sh
chmod +x /path/to/your/go/script/hello.cgi
```

#### 5. Restart NGINX

Restart NGINX to apply the new configuration:

```sh
sudo systemctl restart nginx
```

### Testing the Setup

With the Go server running via `fcgiwrap` and NGINX configured to use FastCGI, you should be able to make a request to the NGINX server on port 80, which will proxy the request to the Go script.

Open your browser and navigate to `http://localhost`. You should see "Hello, World!" as the response.

### Notes

- While using FastCGI with Go is possible, it’s not a common practice. Go applications are usually compiled to standalone binaries and run independently, often behind a reverse proxy like NGINX using the proxy_pass directive as shown in the first example.
- The `fcgiwrap` method is more common for scripting languages like PHP. For Go, it’s recommended to use the standard method of running the Go binary and using `proxy_pass` in NGINX for better performance and simplicity.