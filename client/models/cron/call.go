package cron

import (
	"fmt"

	. "github.com/yunkaiyueming/netburn/client/service"
	. "github.com/yunkaiyueming/netburn/g"
	"github.com/yunkaiyueming/netburn/utils"
)

type CronClient struct{}

var DefaultCronClient = new(CronClient)

func (c *CronClient) GetAllCrons() []CronJob {
	var reply []CronJob
	err := RpcClient.Call("CronServer.AllCron", nil, &reply)
	utils.LogErr("models/cron/GetAllCrons", err)

	for _, v := range reply {
		fmt.Println(v)
	}
	return reply
}
