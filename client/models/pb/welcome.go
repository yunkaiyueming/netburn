package pb

import (
	"fmt"
	"log"

	"github.com/yunkaiyueming/netburn/client/service"
	"github.com/yunkaiyueming/netburn/g/welcome"
	"golang.org/x/net/context"
)

func WelCall() string {
	c := welcome.NewWelcomeClient(service.GrpcClient)
	name := "word"

	r, err := c.SayWel(context.Background(), &welcome.FromRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("welcome rpc return : %s,%s,%s", r.Messages, r.GetWe().GetSourth(), r.GetWe().GetDate())
	return fmt.Sprintf("welcome rpc return : %s,%s,%s", r.Messages, r.GetWe().GetSourth(), r.GetWe().GetDate())
}
