package db

import (
	"fmt"

	"github.com/yunkaiyueming/netburn/utils"
	"gopkg.in/mgo.v2"
)

var MgoSession *mgo.Session

func getMongoConf() string {
	mgoCf, err := utils.GetMgoConfig()
	if err != nil {
		fmt.Println("mgo cfg err:" + err.Error())
		panic(err)
	}
	return mgoCf["mongo"]
}

func initMgoConn() {
	session, err := mgo.Dial(getMongoConf())
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	MgoSession = session
}

func GetAuthDb() *mgo.Database {
	return GetDbByName("yama_admin")
}

func GetDbByName(dbName string) *mgo.Database {
	return MgoSession.DB(dbName)
}
