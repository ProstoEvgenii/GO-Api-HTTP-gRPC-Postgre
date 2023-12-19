package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func requestHandler(ctx *fasthttp.RequestCtx) {
	// if fmt.Fprintf("%v", ctx.RequestURI()) == "api" {

	// }
	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
}

func loadConfig() {
	// Установка значений по умолчанию для конфигурационных параметров
	viper.SetDefault("service_a.port", 8080)
	// Загрузка файлов конфигурации .env
	if err := godotenv.Load(); err != nil {
		log.Println("Ошибка загрузки .env")
	}
	// Чтение переменных окружения
	viper.AutomaticEnv()
}

func createServiceBConnection() *grpc.ClientConn {
	// Получение адреса сервера ServiceB из конфигурации
	serviceBAddress := viper.GetString("service_b.address")
	creds := credentials.NewTLS(nil)
	// Установка параметров подключения к gRPC серверу ServiceB
	conn, err := grpc.Dial(serviceBAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("Failed to connect to ServiceB: %v", err)
	}
	return conn
}
