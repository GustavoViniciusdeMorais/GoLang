module example.com/api

go 1.18

replace example.com/cmds => ../cmds

replace example.com/handlers => ../handlers

require (
	example.com/cmds v0.0.0-00010101000000-000000000000
	example.com/handlers v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)
