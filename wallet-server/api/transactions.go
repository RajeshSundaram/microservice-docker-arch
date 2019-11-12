package api

import (
	"encoding/json"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([2]string{"BTC", "ETH"})
}
