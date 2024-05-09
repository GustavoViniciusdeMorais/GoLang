package src

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET Request\n")
}

func (a *App) Initialize() {
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	a.DB = db
	a.Port = ":9003"
	a.Router = mux.NewRouter()
}

func (a *App) Run() {
	a.Router.HandleFunc("/", getRequest).Methods("GET")

	fmt.Printf("Server at localhost:%v\n", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
