package cron

import (
	"fmt"
	"time"
)

func CronTest1() {
	CreateCronJob(1*time.Minute, test1, "CronTest1")
}

func CronTest2() {
	CreateCronJob(20*time.Second, test2, "CronTest2")
}

func CronTest3() {
	CreateCronJob(25*time.Second, test3, "CronTest3")
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
