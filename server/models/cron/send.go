package cron

import (
	"fmt"

	. "github.com/yunkaiyueming/netburn/g"
	"github.com/yunkaiyueming/netburn/server/cron"
	"github.com/yunkaiyueming/netburn/utils"
)

type CronServer struct{}

func (s *CronServer) AllCron(input interface{}, jobs *[]CronJob) error {
	utils.LogCall("CronServer:AllCron", "", "ok")

	rets := cron.GetSortedCron()
	for _, ret := range rets {
		tmp := CronJob{
			Name:          ret.Name,
			IsRun:         ret.IsRun,
			NextStartTime: ret.NextStartTime,
		}
		*jobs = append(*jobs, tmp)
	}
	fmt.Println(jobs)
	return nil
}
