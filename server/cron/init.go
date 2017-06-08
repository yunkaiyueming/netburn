package cron

import (
	"fmt"
	"time"

	"runtime"

	"github.com/astaxie/beego/logs"
)

func init() {
}

func CreateCronJob(d time.Duration, f func() string) {
	Ticker := time.NewTicker(d)
	runLock := false

	for {
		t := <-Ticker.C
		if !runLock {
			runLock = true

			pc, _, _, _ := runtime.Caller(1)
			fName := runtime.FuncForPC(pc).Name()
			logs.Alert(fmt.Sprintf("fname:%v,start_time:%s\n", fName, t.Format("2006-01-02 15:04:05")))
			ret := f()
			fmt.Println(ret)
			logs.Alert(fmt.Sprintf("fanme:%v,end_time:%s,ret:%s\n", fName, t.Format("2006-01-02 15:04:05")), ret)

			runLock = false
		}
	}
}
