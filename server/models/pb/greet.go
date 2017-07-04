package pb

import (
	pb "github.com/yunkaiyueming/netburn/g"
	"golang.org/x/net/context"
)

// server is used to implement helloworld.GreeterServer.
type GreetServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *GreetServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "grpc Hello " + in.Name}, nil
}
