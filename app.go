package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func fetchPrice(ticker string) (string, error) {
	url := "https://api.iextrading.com/1.0/stock/" + ticker + "/price"	

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	price, err := fetchPrice("aapl")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(price)
}
