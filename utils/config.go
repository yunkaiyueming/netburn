package utils

import (
	"github.com/astaxie/beego/config"
)

func GetServerConfig() (map[string]string, error) {
	cf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		return nil, err
	}
	return cf.GetSection("server")
}
