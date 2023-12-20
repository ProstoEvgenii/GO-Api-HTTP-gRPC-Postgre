package main

import (
	"context"
	"fmt"
	"log"
	"service_1/proto"

	"github.com/joho/godotenv"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	host := "127.0.0.1:888"
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Файл ENV не найден")
		host = ":888"
	}

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	router := routing.New()

	client := proto.NewService1Client(conn)

	router.Get("/api", func(ctx *routing.Context) error {

		symbol := string(ctx.QueryArgs().Peek("symbol"))
		interval := string(ctx.QueryArgs().Peek("interval"))

		resp, _ := client.GetDataFromApi(context.TODO(), &proto.Request{
			Symbol:   symbol,
			Interval: interval,
		})
		fmt.Fprintf(ctx, "Response: %s", resp.DataBinance)
		return nil
	})
	fasthttp.ListenAndServe(host, router.HandleRequest)

}
