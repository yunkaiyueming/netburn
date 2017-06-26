package cache

import (
	//"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/yunkaiyueming/netburn/server/db"
)

type CacheServer struct{}

func (c *CacheServer) GetHotData(input interface{}, output *interface{}) error {
	db.RsCon.Do("SELECT", 0)
	db.RsCon.Do("set", "homework", "big hourse big table")
	v, err := redis.String(db.RsCon.Do("get", "homework"))
	*output = v

	return err
}
