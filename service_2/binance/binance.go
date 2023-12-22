package binance

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type BinanceData struct {
	Timestamp int64  `json:"timestamp"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	// CloseTimestamp        int64
	// QuoteAssetVolume      string
	// NumberOfTrades        int
	// TakerBuyBaseAssetVol  string
	// TakerBuyQuoteAssetVol string
	// Ignore                string
}

func GetDataFromBinance(symbol, interval string) ([]byte, error) {
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
