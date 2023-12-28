package backend

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

func (a *App) allProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(a.DB)
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
	err := p.getProduct(a.DB)
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

	err := p.createProduct(a.DB)
	if err != nil {
		fmt.Printf("createProduct error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, p)
}

func (a *App) allOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := getOrders(a.DB)
	if err != nil {
		fmt.Printf("getOrders error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, orders)
}

func (a *App) fetchOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var order Order
	order.Id, _ = strconv.Atoi(id)
	err := order.getOrder(a.DB)
	if err != nil {
		fmt.Printf("fetchOrder error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusOK, order)
}

func (a *App) createOrder(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var order Order
	json.Unmarshal(reqBody, &order)

	err := order.createOrder(a.DB)
	if err != nil {
		fmt.Printf("createOrder error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	for _, item := range order.Items {
		var orderItem OrderItem
		orderItem = item
		orderItem.OrderId = order.Id
		err := orderItem.createOrderItem(a.DB)
		if err != nil {
			fmt.Printf("createOrder error: %s\n", err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}

	respondWithJson(w, http.StatusOK, order)
}

func (a *App) Run() {
	a.Router.HandleFunc("/", getRequest).Methods("GET")
	a.Router.HandleFunc("/products", a.allProducts).Methods("GET")
	a.Router.HandleFunc("/products/{id}", a.fetchProduct).Methods("GET")
	a.Router.HandleFunc("/products", a.createProduct).Methods("POST")

	a.Router.HandleFunc("/orders", a.allOrders).Methods("GET")
	a.Router.HandleFunc("/orders/{id}", a.fetchOrder).Methods("GET")

	fmt.Printf("Server at localhost:%v\n", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
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
