package slog

import (
	//	"fmt"

	. "github.com/yunkaiyueming/netburn/client/service"
	"github.com/yunkaiyueming/netburn/utils"
)

type SlogClient struct{}

var DefaultSlogClient = new(SlogClient)

func (s *SlogClient) GetApps() interface{} {
	var ret interface{}
	err := RpcClient.Call("SlogServer.GetApps", nil, &ret)
	utils.LogErr("SlogClient/GetApps", err)
	return ret
}

func (s *SlogClient) GetAppByName(appName string) interface{} {
	var ret interface{}
	err := RpcClient.Call("SlogServer.GetAppByName", appName, &ret)
	utils.LogErr("SlogClient/GetApps", err)
	return ret
}

func (s *SlogClient) GetAppLog(appName string) interface{} {
	var ret interface{}
	err := RpcClient.Call("SlogServer.GetAppLog", appName, &ret)
	utils.LogErr("SlogClient/GetApps", err)
	return ret
}
