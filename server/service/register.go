package service

import (
	//"fmt"
	"net/rpc"

	"github.com/yunkaiyueming/netburn/server/models/users"
)

func init() {
	rpc.Register(new(users.UserServer))

	startService()
}
