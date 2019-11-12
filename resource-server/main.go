package main

import (
	"net/http"

	"./api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json")

	// Transactions
	trx := router.PathPrefix("/trx").Subrouter()
	trx.HandleFunc("/list", api.GetTransactions)

	http.ListenAndServe(":8080", router)
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello User"))
}
