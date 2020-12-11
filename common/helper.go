package common

import (
	"net/http"
	"strings"

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
type AjaxResponseData struct {
	Status  int
	Message string
	Content interface{}
}

func Ajax(status int, message string, content interface{}) (data AjaxResponseData) {
	data.Status = status
	data.Message = message
	data.Content = content
	return data
}

// 生成完整url
func GetUrl(r *http.Request, uri string) string {
	referer := r.Referer()
	sep := "://"
	protocol := strings.Split(referer, sep)[0]
	url := protocol + sep + r.Host + uri
	return url
}
