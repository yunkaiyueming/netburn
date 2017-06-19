package cron

import (
	"fmt"
	"time"

	"github.com/yunkaiyueming/netburn/server/db"
	"github.com/yunkaiyueming/netburn/server/models/users"
)

func CronUserLand() {
	CreateCronJob(30*time.Second, userLand, "CronUserLand")
}

func userLand() string {
	items := make([]map[string]interface{}, 0)
	uInfos := users.GetDefaultUsers()
	for _, u := range uInfos {
		item := map[string]interface{}{
			"email": u.Email,
			"name":  u.Name,
			"age":   u.Age,
			"memo":  u.Memo,
		}

		items = append(items, item)
	}

	updateKeys := []string{"name", "age", "memo"}
	baseModel := db.ModelBase{10, []string{"email"}, db.UserFieldsDefaultValue}
	ret := baseModel.UpsertBatch("default", db.USER_TABLE_NAME, items, updateKeys)
	return fmt.Sprintf("ret==>%s:%t", "userLand/UpsertBatch", ret)
}
