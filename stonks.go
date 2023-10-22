package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

//go:embed secret.key
var API_KEY string

const POLYGON_PATH = "https://api.polygon.io"

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
	Price  float64
}

type Values struct {
	Open float64 `json:"open"`
}

func SearchTicker(ticker string) []Stock {
	// resp, err := http.Get("https://finnhub.io/api/v1/search?q=" + key + "&token=" + os.Getenv("FINNHUB_API_KEY"))
	// https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/2023-01-09/2023-01-09?apiKey=V7dKvb6Az7cN75VdYDbluMa0YwJniMJu
	resp, err := http.Get(POLYGON_PATH + "/v2/aggs/ticker/" + strings.ToUpper(ticker) + "?apiKey=" + API_KEY)
	if err != nil {
		log.Fatal(err)
	}
	// defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	data := struct {
		Results []Stock `json:"results"`
	}{}

	json.Unmarshal(body, &data)

	return data.Results
}
