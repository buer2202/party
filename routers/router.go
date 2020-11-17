package routers

import (
	"party2202.com/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 主页
	beego.Router("/party/:id", &controllers.PartyController{})
}
