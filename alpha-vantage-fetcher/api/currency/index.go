package currency

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	api ".."
	"../../infra/rest"
	"../../util/csvutil"
)

type Currency struct {
	Name string `json:"code"`
	Code string `json:"name"`
	Type string `json:"type"`
}

type CurrencyType struct {
	Type     string
	SitePath string
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
