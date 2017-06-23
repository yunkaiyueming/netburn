package slog

import (
	//"fmt"

	"github.com/yunkaiyueming/netburn/server/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SlogServer struct{}

func (s *SlogServer) getCol() *mgo.Collection {
	return db.GetAuthDb().C("app")
}

func (s *SlogServer) getLogCol(appName string) *mgo.Collection {
	return db.GetDbByName("yama_" + appName).C("url_view")
}

func (s *SlogServer) GetApps(input interface{}, output *[]interface{}) error {
	return s.getCol().Find(nil).All(output)
}

func (s *SlogServer) GetAppByName(appName interface{}, output *interface{}) error {
	return s.getCol().Find(bson.M{"app_name": appName}).One(output)
}

func (s *SlogServer) GetAppLog(appName interface{}, output *[]interface{}) error {
	return s.getLogCol(appName.(string)).Find(nil).All(output)
}
