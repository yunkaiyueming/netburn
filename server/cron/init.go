package cron

import (
	"fmt"
	"time"

	"runtime"

	"github.com/astaxie/beego/logs"
	"github.com/yunkaiyueming/netburn/server/safe"
)

var jobs safe.SafeJobList

func init() {
	go CronUserLand()
	go CronTest1()
	go CronTest2()
	go CronTest3()
}

func CreateCronJob(d time.Duration, f func() string, name string) {
	Ticker := time.NewTicker(d)
	runLock := false

	d.Seconds()
	newJob := safe.CronJob{
		Name:          name,
		IsRun:         "false",
		NextStartTime: time.Unix(time.Now().Unix()+int64(d.Seconds()), 0).Format("2006-01-02 15:04:05"),
	}
	jobs.AddJob(newJob)

	pc, _, _, _ := runtime.Caller(1)
	fName := runtime.FuncForPC(pc).Name()
	for {
		<-Ticker.C
		if !runLock {
			runLock = true

			logs.Alert(fmt.Sprintf("fname:%v,start_time:%s\n", fName, time.Now().Format("2006-01-02 15:04:05")))

			nextStarTime := time.Unix(time.Now().Unix()+int64(d.Seconds()), 0).Format("2006-01-02 15:04:05")
			jobs.UpdateCronByName(name, nextStarTime, "true")
			ret := f()
			jobs.UpdateCronByName(name, nextStarTime, "false")

			logs.Alert(fmt.Sprintf("fanme:%v,end_time:%s,ret:%s\n", fName, time.Now().Format("2006-01-02 15:04:05"), ret))
			runLock = false
		}
	}
}

func GetSortedCron() []safe.CronJob {
	return jobs.GetSortedJobList()
}
