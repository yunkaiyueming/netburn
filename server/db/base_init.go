package db

import (
	"fmt"
)

func init() {
	err := initORM()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	//mgo conn
	//MgoSession = GetDbConn()

	//redis conn

}
