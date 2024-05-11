package src

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
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

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET Request\n")
}

func (a *App) allProducts(w http.ResponseWriter, r *http.Request) {
	products, err := GetProducts(a.DB)
	if err != nil {
		fmt.Printf("getProducts error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, products)
}

func (a *App) fetchProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var p Product
	p.Id, _ = strconv.Atoi(id)
	err := p.GetProduct(a.DB)
	if err != nil {
		fmt.Printf("fetchProduct error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, p)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var p Product
	json.Unmarshal(reqBody, &p)

	err := p.CreateProduct(a.DB)
	if err != nil {
		fmt.Printf("createProduct error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, p)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) Run() {
	a.Router.HandleFunc("/", getRequest).Methods("GET")
	a.Router.HandleFunc("/products", a.allProducts).Methods("GET")
	a.Router.HandleFunc("/products/{id}", a.fetchProduct).Methods("GET")
	a.Router.HandleFunc("/products", a.createProduct).Methods("POST")

	fmt.Printf("Server at localhost:%v\n", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
