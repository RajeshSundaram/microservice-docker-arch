package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
)

func GetResponseAsType(url string, bodyType interface{}) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("The call to " + url + " is failed")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error occured whild parsing " + url + " response with " + reflect.TypeOf(bodyType).Name())
	}
	err = json.Unmarshal(body, &bodyType)
	if err != nil {
		return nil, errors.New("Unmarshalling error")
	}
	return bodyType, nil
}

func GetResponseAsBytes(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("The call to " + url + " is failed")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error occured whild parsing " + url + " response")
	}
	return body, nil
}
