# GoLang

Created by: Gustavo Morais

[GoLang Certificate](https://www.linkedin.com/learning/certificates/f75363e5c176cbb7695329ac0f68aebb6997b8dc520d0490e53eb440f67fe548)

### Libraries
- [HTTP Router](https://github.com/gorilla/mux)
  - go install github.com/gorilla/mux@latest
- [SQLite](https://github.com/mattn/go-sqlite3)
  - go install github.com/mattn/go-sqlite3@latest

### Build module doc example
- [Module Doc](./CreatingModule.md)

### Troble shooting go commands
```sh
export PATH=$PATH:/usr/local/go/bin
```

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