package common

import (
	"github.com/astaxie/beego"
)

// MyLog 日志
func MyLog(fileName string, contents string) {
	beego.SetLogger("file", `{"filename":"logs/`+fileName+`.log","daily":true}`)
	beego.BeeLogger.DelLogger("console")
	beego.SetLogFuncCall(true)
	beego.Notice(contents)
}

// ajax响应格式
type ajax struct {
    Status  int
    Message string
    Content string
}
func Ajax(status int, message string, content string) (data ajax) {
	data.Status = status
	data.Message = message
	data.Content = content
	return data
}
