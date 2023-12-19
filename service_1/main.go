package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	host := "127.0.0.1:8080"
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Файл ENV не найден")
		host = ":8080"
	}

	// serviceBConn := createServiceBConnection()
	// defer serviceBConn.Close()
	// loadConfig()
	router := routing.New()

	router.Get("/api/<data>", func(ctx *routing.Context) error {
		log.Printf("Name: %v", ctx.Param("data"))
		fmt.Fprintf(ctx, "Name: %v", ctx.Param("data"))
		return nil
	})
	fasthttp.ListenAndServe(host, router.HandleRequest)

}

// func createServiceBConnection() *grpc.ClientConn {
// 	// Получение адреса сервера ServiceB из конфигурации
// 	serviceBAddress := viper.GetString("service_b.address")
// 	creds := credentials.NewTLS(nil)
// 	// Установка параметров подключения к gRPC серверу ServiceB
// 	conn, err := grpc.Dial(serviceBAddress, grpc.WithTransportCredentials(creds))
// 	if err != nil {
// 		log.Printf("Failed to connect to ServiceB: %v", err)
// 	}
// 	return conn
// }
