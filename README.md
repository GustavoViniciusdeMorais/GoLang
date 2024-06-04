# DDD API

```sh
```

### Libraries
- [HTTP Router](https://github.com/gorilla/mux)
  - go get github.com/gorilla/mux@latest
- [GoLang SQLite](https://github.com/mattn/go-sqlite3)
  - go get github.com/mattn/go-sqlite3@latest
- [SQLite](https://www.sqlite.org/docs.html)
	- sudo apt install sqlite3
	- [Sqlite3 Doc](https://github.com/GustavoViniciusdeMorais/Database_Studies/tree/sqlite)

### JWT with Redis
- [Jwt with Redis](./auth_jwt_redis.md)
### Build module doc example
- [Module Doc](./CreatingModule.md)
### Database SQLite Example
- [SQLite example](./DB.md)
### Troble shooting go commands
```sh
export PATH=$PATH:/usr/local/go/bin
```
### API Test
- [API](./api.md)
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

```sh
cd src
go mod init
go get github.com/gorilla/mux@latest
go get github.com/mattn/go-sqlite3@latest
```
