package db

import (
	"gopkg.in/mgo.v2"
)

var MgoSession *mgo.Session

func getMongoConf() string {
	return ""
}

func GetDbConn() *mgo.Session {
	session, err := mgo.Dial(getMongoConf())
	panic(err)

	session.SetMode(mgo.Monotonic, true)
	return session
}

func GetAuthDb() *mgo.Database {
	return GetDbByName("yama_admin")
}

func GetDbByName(dbName string) *mgo.Database {
	return MgoSession.DB(dbName)
}

func GetAuditDb(appName string) *mgo.Database {
	return GetDbByName("yama_log_" + appName)
}
