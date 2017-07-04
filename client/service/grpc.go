package service

import (
	"log"

	"google.golang.org/grpc"
)

var GrpcClient *grpc.ClientConn

func StartGrpc() {
	var err error
	GrpcClient, err = grpc.Dial("127.0.0.1:8093", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer GrpcClient.Close()
}
