package routers

import (
	"mybeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 主页
	beego.Router("/party", &controllers.PartyController{})
}
