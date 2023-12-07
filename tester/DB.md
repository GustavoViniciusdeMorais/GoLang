# Go Sqlite

```sh
sudo go run configDB.go
sudo go run execDbCommands.go
```
### Connection example
```go
func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SQLite version ", version)
	return db, err
}
```