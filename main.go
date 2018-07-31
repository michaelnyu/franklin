package main

import (
	"flag"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
)

const (
	apibase = "https://api.iextrading.com/1.0"
)

var (
	ticker string
	detail bool
)

// function that fetches the current price of the ticker
func fetchPrice(ticker string) (string, error) {
	url := fmt.Sprintf("%s/stock/%s/price", apibase, ticker)

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

// fetches the detailed stats on a ticker
func fetchDetails(ticker string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/stock/%s/stats", apibase, ticker)
	var data map[string]interface{}

	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	// unmarshall response body
    err = json.Unmarshal(body, &data)
    if err != nil {
        panic(err)
    }

	return data, nil
}

// filters selected parts from the fetchDetails response
func parseDetails(data map[string]interface{}) (string, error) {
	var res string

	fields := []string{"week52high", "week52low", "dividendRate", "dividendYield", "beta", "returnOnEquity", "day200MovingAvg", "day50MovingAvg"}
	err := errors.New("Error parsing details")

	if val, ok := data["companyName"]; ok {
		res = fmt.Sprintf("\t%s has a ...", val)
	} else {
		return "", err
	}

	for _, field := range fields {
		if val, ok := data[field]; ok {
			res += fmt.Sprintf("\n\t  %s of %f", field, val)
		} else {
			return "", err
		}
	}

	return res, nil
}

func init() {
    flag.StringVar(&ticker, "ticker", "", "ticker to get data of")
	flag.StringVar(&ticker, "t", "", "(alias) ticker to get data of")
	flag.BoolVar(&detail, "detail", false, "show details")
	flag.BoolVar(&detail, "d", false, "(alias) show details")

    flag.Usage = func() {
        flag.PrintDefaults()
    }

	flag.Parse()
}

func main() {
	fmt.Println("💵")

	if (detail) {
		details, err := fetchDetails(ticker)
		if err != nil {
			log.Fatal(err)
		}
		res, err := parseDetails(details)
		fmt.Println(res)
	} else {
		price, err := fetchPrice(ticker)
		if err != nil {
			log.Fatal(err)
		}
		out := fmt.Sprintf("\tThe price of %s currently is: %s USD", ticker, price)
		fmt.Println(out)
	}

	fmt.Println("💵")
}
