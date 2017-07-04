package pb

import (
	"time"

	"github.com/yunkaiyueming/netburn/g/welcome"
	"golang.org/x/net/context"
)

type WelcomeServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *WelcomeServer) SayWel(ctx context.Context, in *welcome.FromRequest) (*welcome.ToReply, error) {
	return &welcome.ToReply{
		"welcome " + in.Name,
		&welcome.Weather{"北京", time.Now().Format("2006-01-02")},
	}, nil
}
