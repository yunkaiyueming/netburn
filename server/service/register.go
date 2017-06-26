package service

import (
	//"fmt"
	"net/rpc"

	"github.com/yunkaiyueming/netburn/server/models/cache"
	"github.com/yunkaiyueming/netburn/server/models/cron"
	"github.com/yunkaiyueming/netburn/server/models/hfile"
	"github.com/yunkaiyueming/netburn/server/models/slog"
	"github.com/yunkaiyueming/netburn/server/models/users"
)

func init() {
	//first register
	rpc.Register(new(users.UserServer))
	rpc.Register(new(cron.CronServer))
	rpc.Register(new(hfile.HfileServer))
	rpc.Register(new(slog.SlogServer))
	rpc.Register(new(cache.CacheServer))
}
