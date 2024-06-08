# GoLang

Created by: Gustavo Morais

[GoLang Certificate](https://www.linkedin.com/learning/certificates/f75363e5c176cbb7695329ac0f68aebb6997b8dc520d0490e53eb440f67fe548)

### Troble shooting go commands
```sh
export PATH=$PATH:/usr/local/go/bin
```

### Tutorials
- [NGINX Proxy](./goWithNginx.md)
- [NGINX CGI](./goWithCgiNginx.md)

### Other go installation
```sh
sudo apt install golang-go -y
go clean -modcache
go install -v golang.org/x/tools/gopls@latest
```

### Config go env
Type the following command to start the Go workspace.
This fixes the error of importing packages
```sh
cd project/path/
go work init
```

```sh
go version
go run first.go
```

### Simple example
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func gus(name string) {
	fmt.Printf("Entered text was %s", name)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	input, err := reader.ReadString('\n')

	if err != nil {
		panic("error")
	}

	gus(input)
}
```