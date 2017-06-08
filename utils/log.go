package utils

import (
	_ "fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func LogErr(title string, err error) {
	if err != nil {
		beego.SetLogger("multifile", `{"filename":"./log/err.log","level":3,"separate":["error"]}`)
		logs.EnableFuncCallDepth(true)
		logs.SetLogFuncCallDepth(1)
		logs.Error(title, err.Error())
	}
}

func LogInfo(title, content string) {
	beego.SetLogger("multifile", `{"filename":"./log/info.log","level":4,"separate":["info"]}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(1)
	logs.Info(title, content)
}

func LogCall(fname, args, ret string) {
	beego.SetLogger("multifile", `{"filename":"./log/call.log","level":4,"separate":["info"]}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(6)
	logs.Info("===>%s(%s),<===%s", fname, args, ret)
}
