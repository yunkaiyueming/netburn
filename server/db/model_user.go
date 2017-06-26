package db

import (
	"fmt"

	_ "github.com/astaxie/beego/orm"
)

type UserModel struct{}

const USER_TABLE_NAME = "user"

var UserFieldsDefaultValue = []map[string]interface{}{
	{"name": ""},
	{"age": 0},
	{"memo": ""},
}

func (u *UserModel) UpdateUserByEmail(table, email string, data map[string]interface{}) error {
	sql := ""
	for k, v := range data {
		sql += fmt.Sprintf("%s='%s',", k, v.(string))
	}

	sql = fmt.Sprintf("UPDATE %s SET %s", table, sql)
	updateSql := sql[0 : len(sql)-1]
	fmt.Println(updateSql)
	_, err := GetOrm("default").Raw(updateSql).Exec()
	return err
}
