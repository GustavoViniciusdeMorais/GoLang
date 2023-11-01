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

```sh
go version
go run first.go
```