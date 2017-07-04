package service

import (
	"fmt"
	"log"
	"net"

	"github.com/yunkaiyueming/netburn/g"
	"github.com/yunkaiyueming/netburn/g/welcome"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/yunkaiyueming/netburn/server/models/pb"
)

const (
	port = "127.0.0.1:8093"
)

var s *grpc.Server

func StartGrpc() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Println("grpc server start ===>", port)
	}
	s = grpc.NewServer()

	//register server
	g.RegisterGreeterServer(s, &pb.GreetServer{})
	welcome.RegisterWelcomeServer(s, &pb.WelcomeServer{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GetGServerInfo() string {
	var info string
	for k, v := range s.GetServiceInfo() {
		s := ""
		for _, m := range v.Methods {
			s += fmt.Sprintf("%s:%t-%t\t", m.Name, m.IsClientStream, m.IsServerStream)
		}
		info += fmt.Sprintf("%s,%s,%s\n", k, v.Metadata, s)
	}
	return info
}
