package main

import (
	"flag"
	"fmt"
	"net/http"
	"io/ioutil"
)

var (
	ticker string
	detail bool
)

func fetchPrice(ticker string) (string, error) {
	url := fmt.Sprintf("https://api.iextrading.com/1.0/stock/%s/price", ticker)

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

func init() {
    flag.StringVar(&ticker, "ticker", "", "ticker to get data of")
	flag.StringVar(&ticker, "t", "", "(alias) ticker to get data of")
	flat.BoolVar(&)

    flag.Usage = func() {
        flag.PrintDefaults()
    }

	flag.Parse()
}

func main() {
	price, err := fetchPrice(ticker)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(price)
}
