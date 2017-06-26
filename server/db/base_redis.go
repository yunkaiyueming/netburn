package db

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/yunkaiyueming/netburn/utils"
)

var RsCon redis.Conn

func initRedisConn() {
	cf, err := utils.GetRedisConf()
	if err != nil {
		panic(err)
		fmt.Println("get redis config error:", err.Error())
	}

	//dialoption := redis.DialPassword("admin")
	//rs, err = redis.Dial("tcp", "localhost:6379", dialoption)
	RsCon, err = redis.DialURL(cf["redis"])
	if err != nil {
		fmt.Println("redis connect error", err.Error())
		panic(err)
	}

}

func CloseRedisConn() {
	RsCon.Close()
}

func GetRedisCon() redis.Conn {
	if RsCon == nil {
		initRedisConn()
	}
	return RsCon
}
