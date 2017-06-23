package hfile

import (
	"fmt"

	. "github.com/yunkaiyueming/netburn/client/service"
	//. "github.com/yunkaiyueming/netburn/g"
	"github.com/yunkaiyueming/netburn/utils"
)

type HFileClient struct{}

var DefaultHfileClient = new(HFileClient)

func (h *HFileClient) AllBucketsFiles() interface{} {
	var reply interface{}
	err := RpcClient.Call("HfielServer.GetAllBucketsFiles", nil, &reply)
	utils.LogErr("HFileClient/AllBucketsFiles", err)
	return reply
}

func (h *HFileClient) GetBucketFiles(bucketName string) interface{} {
	var reply interface{}
	err := RpcClient.Call("HfileServer.GetBucketFiles", bucketName, &reply)
	utils.LogErr("HFileClient/GetBucketFiles", err)
	return reply
}

func (h *HFileClient) GetBucketNames() interface{} {
	var reply interface{}
	err := RpcClient.Call("HfileServer.GetBucketNames", nil, &reply)
	utils.LogErr("HFileClient/GetBucketNames", err)
	return reply
}

func (h *HFileClient) GetFileMsg(bucketName, key string) interface{} {
	var reply interface{}
	input := map[string]string{
		"bucket_name": bucketName,
		"key":         key,
	}
	err := RpcClient.Call("HfileServer.GetFileMsg", input, &reply)
	utils.LogErr("HFileClient/GetFileMsg", err)
	fmt.Println(input, reply)
	return reply
}
