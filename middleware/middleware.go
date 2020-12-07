package middleware

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/", beego.BeforeRouter, FilterGuest)
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
}
