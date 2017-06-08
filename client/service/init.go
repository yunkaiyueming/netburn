package service

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

var RpcClient *rpc.Client

func init() {
	if RpcClient == nil {
		var err error
		RpcClient, err = jsonrpc.Dial("tcp", "127.0.0.1:8092")
		if err != nil {
			panic(err)
		}
	}
}
