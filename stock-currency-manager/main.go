package main

import (
	"net/http"

	"scm/api"
	"scm/infra"

	"github.com/gorilla/mux"
)

func main() {
	Init()
	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json")

	router.HandleFunc("/env", sayHello).Methods("GET")

	currencyRoute := router.PathPrefix("/currency").Subrouter()
	currencyRoute.HandleFunc("/list", api.GetAllCurrencies)
	currencyRoute.HandleFunc("/exchange/{from}/{to}", api.GetCurrencyExchangeResult)

	// Transactions
	trx := router.PathPrefix("/trx").Subrouter()
	trx.HandleFunc("/list", api.GetTransactions)

	http.ListenAndServe(":8080", router)
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(infra.GetDbConnectionURL()))
}

func Init() {
	InitializeEnvironment()
	InitializeDatabase()
}
