package pb

import (
	"log"

	"github.com/yunkaiyueming/netburn/client/service"
	"github.com/yunkaiyueming/netburn/g"
	"golang.org/x/net/context"
)

func GreetCall() string {
	c := g.NewGreeterClient(service.GrpcClient)
	name := "word"

	r, err := c.SayHello(context.Background(), &g.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	return r.Message
}
