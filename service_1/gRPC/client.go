package grpc

import (
	"context"
	"log"
	"service_1/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Client proto.Service1Client

func Start_gRPCClient() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	Client = proto.NewService1Client(conn)
}

func GetDataFromAPI(symbol, interval string) (*proto.Response, error) {

	ctx := context.Background()

	response, err := Client.GetDataFromApi(ctx, &proto.Request{
		Symbol:   symbol,
		Interval: interval,
	})
	if err != nil {
		return nil, err
	}

	return response, err
}
