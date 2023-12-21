package main

import (
	"log"
	grpc "service_1/gRPC"
	"service_1/server"

	"github.com/joho/godotenv"
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

func main() {
	host := "127.0.0.1:888"
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Файл ENV не найден")
		host = ":888"
	}
	grpc.Start_gRPCClient()
	server.Start(host)

}
