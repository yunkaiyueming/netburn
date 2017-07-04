package cache

import (
	. "github.com/yunkaiyueming/netburn/client/service"
	"github.com/yunkaiyueming/netburn/utils"
)

type CacheClient struct{}

var DefaultCacheClient CacheClient

func (c *CacheClient) GetHotData() interface{} {
	var ret interface{}
	err := RpcClient.Call("CacheServer.GetHotData", nil, &ret)
	utils.LogErr("SlogClient/GetApps", err)
	return ret
}
