package common

import "github.com/astaxie/beego"

// MyLog 日志
func MyLog(fileName string, contents string) {
    beego.SetLogger("file", `{"filename":"logs/`+fileName+`.log","daily":true}`)
    beego.BeeLogger.DelLogger("console")
    beego.SetLogFuncCall(true)
	beego.Notice(contents)
}
