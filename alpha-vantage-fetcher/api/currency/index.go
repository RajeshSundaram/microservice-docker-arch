package currency

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	api ".."
	"../../infra/rest"
	"../../util/csvutil"
)

var alphaVantageApiKey = os.Getenv("ALPHA_VANTAGE_API_KEY")

type Currency struct {
	Name string `json:"code"`
	Code string `json:"name"`
	Type string `json:"type"`
}

type CurrencyType struct {
	Type     string
	SitePath string
}

type CurrencyExchangeResponse struct {
	Result CurrencyExchangeResult `json:"Realtime Currency Exchange Rate"`
}

type CurrencyExchangeResult struct {
	ExchangeRate float64 `json:"5. Exchange Rate"`
	BidPrice     float64 `json:"8. Bid Price"`
	AskPrice     float64 `json:"9. Ask Price"`
}

var digitalCurrencyType = CurrencyType{Type: "Digital", SitePath: "digital_currency_list"}
var physicalCurrencyType = CurrencyType{Type: "Physical", SitePath: "physical_currency_list"}

func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	digitalCurrencies, err := getCurrencies(digitalCurrencyType)
	if err != nil {
		log.Panic(err)
		w.Write([]byte(err.Error()))
		return
	}
	physicalCurrencies, err := getCurrencies(physicalCurrencyType)
	if err != nil {
		log.Panic(err)
		w.Write([]byte(err.Error()))
		return
	}
	currencies := append(physicalCurrencies, digitalCurrencies...)
	resData, _ := json.Marshal(currencies)
	w.Write(resData)
}

func GetCurrencyExchangeResult(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := fmt.Sprintf("%s/query?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=%s&apikey=%s", api.AlphaVantageRootURL, params["from"], params["to"], alphaVantageApiKey)
	res := &CurrencyExchangeResponse{}
	// TODO:
	// res, err := rest.GetResponseAsType(url, res)
	// if err != nil {
	// 	log.Panic(err)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// byteRes, _ := json.Marshal(res.Result)
	// w.Write(byteRes)
}

func getCurrencies(currecyType CurrencyType) ([]Currency, error) {
	query := fmt.Sprintf("%s/%s", api.AlphaVantageRootURL, currecyType.SitePath)
	data, err := rest.GetResponseAsBytes(query)
	if err != nil {
		return []Currency{}, err
	}
	records, err := csvutil.ReadCSV(data)
	if err != nil {
		return []Currency{}, err
	}
	currencies := make([]Currency, 0)
	for _, record := range records[1:] {
		curencyItem := Currency{Code: record[1], Name: record[0], Type: currecyType.Type}
		currencies = append(currencies, curencyItem)
	}
	return currencies, nil
}
