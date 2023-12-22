package main

import (
	"context"
	"log"
	"net"
	"service_2/binance"
	"service_2/pkg/proto"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	proto.UnimplementedService1Server
}

func (s *GRPCServer) GetDataFromApi(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	data, err := binance.GetDataFromBinance(req.Symbol, req.Interval)
	if err != nil {
		return nil, err
	}
	log.Println("data")

	return &proto.Response{
		DataBinance: data,
	}, nil
}
func main() {

	s := grpc.NewServer()
	proto.RegisterService1Server(s, &GRPCServer{})
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("Ошибка net.Listen:", err)
	}
	if err := s.Serve(l); err != nil {
		log.Println("Ошибка gRPC соединения:", err)
	}
}
