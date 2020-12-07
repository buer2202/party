package main

import (
	_ "party2202.com/routers"
	_ "party2202.com/middleware"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
