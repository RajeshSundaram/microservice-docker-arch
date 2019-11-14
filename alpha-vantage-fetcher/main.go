package main

import (
	"net/http"

	"./api"
	"./api/currency"
	"./data"
	"github.com/gorilla/mux"
)

func main() {
	Init()
	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json")

	router.HandleFunc("/env", sayHello).Methods("GET")

	currencyRoute := router.PathPrefix("/currency").Subrouter()
	currencyRoute.HandleFunc("/list", currency.GetAllCurrencies)
	currencyRoute.HandleFunc("/exchange/{from}/{to}", currency.GetCurrencyExchangeResult)

	// Transactions
	trx := router.PathPrefix("/trx").Subrouter()
	trx.HandleFunc("/list", api.GetTransactions)

	http.ListenAndServe(":8080", router)
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(data.GetDbConnectionURL()))
}

func Init() {
	InitializeEnvironment()
	InitializeDatabase()
}
