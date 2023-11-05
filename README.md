# GoLang

Created by: Gustavo Morais

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

### Build go mod to import at other file
```sh
cd directory
go mot init github.com/GustavoViniciusdeMorais/directory
```

```sh
go version
go run first.go
```

### Talk is easy, show me the code
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