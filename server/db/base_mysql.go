package db

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func initORM() (err error) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err = registerDB("default"); err != nil {
		return
	}

	if err = initAllDB(); err != nil {
		return
	}

	return nil
}

//注册其他所有库
func initAllDB() error {
	return nil
}

func registerDB(name string) error {
	cf, err := config.NewConfig("ini", "./conf/db.conf")
	if err != nil {
		return err
	}

	mysqlCf, err := cf.GetSection("mysql")
	if err != nil {
		return err
	}

	if _, ok := mysqlCf[name]; !ok {
		panic(name + "database conf is not exist")
	}
	orm.RegisterDataBase(name, "mysql", mysqlCf[name])
	return nil
}

func getOrm(name string) orm.Ormer {
	o := orm.NewOrm()
	o.Using(name)
	return o
}
