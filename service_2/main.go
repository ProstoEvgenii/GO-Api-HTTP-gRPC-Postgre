package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func getDataFromBinance(symbol, interval string) ([]byte, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/klines?symbol=%s&interval=%s", symbol, interval)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	body := resp.Body()
	return body, nil
}

func main() {
	symbol := "BTCUSDT"
	interval := "1h"

	data, err := getDataFromBinance(symbol, interval)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}
