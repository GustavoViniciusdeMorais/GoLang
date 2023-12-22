module example.com/tester

go 1.18

replace example.com/practice => ../practice

require (
	example.com/backend v0.0.0-00010101000000-000000000000
	example.com/practice v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.18 // indirect
)

replace example.com/backend => ../src/backend
