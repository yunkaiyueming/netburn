package cron

import (
	"fmt"
	"time"
)

func CronTest1() {
	CreateCronJob(5*time.Minute, test1, "CronTest1")
}

func CronTest2() {
	CreateCronJob(15*time.Minute, test2, "CronTest2")
}

func CronTest3() {
	CreateCronJob(10*time.Minute, test3, "CronTest3")
}

func test1() string {
	return fmt.Sprintf("test1 cron run")
}

func test2() string {
	return fmt.Sprintf("test2 cron run")
}

func test3() string {
	return fmt.Sprintf("test3 cron run")
}
