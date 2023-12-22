package server

import (
	"encoding/json"
	"fmt"
	"log"
	grpc "service_1/gRPC"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Data struct {
	Timestamp           int64   `json:"timestamp"`
	Open                string  `json:"open"`
	High                string  `json:"high"`
	Low                 string  `json:"low"`
	Close               string  `json:"close"`
	Volume              string  `json:"volume"`
	QuoteVolume         float64 `json:"quoteVolume"`
	NumberOfTrades      int     `json:"numberOfTrades"`
	TakerBuyBaseVolume  string  `json:"takerBuyBaseVolume"`
	TakerBuyQuoteVolume string  `json:"takerBuyQuoteVolume"`
}

type APIResponse struct {
	Symbol   string `json:"symbol"`
	Interval string `json:"interval"`
	Data     []Data `json:"data"`
}
type BinanceData struct {
	Timestamp             int64
	Open                  string
	High                  string
	Low                   string
	Close                 string
	Volume                string
	CloseTimestamp        int64
	QuoteAssetVolume      string
	NumberOfTrades        int
	TakerBuyBaseAssetVol  string
	TakerBuyQuoteAssetVol string
	Ignore                string
}

func Start(host string) {
	router := routing.New()

	router.Get("/api", func(ctx *routing.Context) error {

		symbol := string(ctx.QueryArgs().Peek("symbol"))
		interval := string(ctx.QueryArgs().Peek("interval"))
		resp, err := grpc.GetDataFromAPI(symbol, interval)
		if err != nil {
			log.Println("GetDataFromAPI Err:", err)
		}
		var binanceData [][]string
		if err := json.Unmarshal(resp.DataBinance, &binanceData); err != nil {
			log.Printf("Ошибка при преобразовании ответа в структуру: %v", err)
		}
		log.Println("11111", binanceData[0])

		// _, errUpload := uploadBinanceData(resp.DataBinance, symbol, interval)
		// if errUpload != nil {
		// 	log.Println("uploadBinanceData Err:", err)
		// }

		fmt.Fprintf(ctx, "Response: %s", resp.DataBinance)
		return nil
	})
	fasthttp.ListenAndServe(host, router.HandleRequest)
}

func uploadBinanceData(binanceBytes []byte, symbol, interval string) (*APIResponse, error) {

	var binanceData APIResponse
	if err := json.Unmarshal(binanceBytes, &binanceData); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Ошибка десериализации: %s", err)
	}

	binanceData.Symbol = symbol
	binanceData.Interval = interval
	log.Println(binanceData)

	return &binanceData, nil
}
