package routers

import (
	"github.com/astaxie/beego"
	"party2202.com/controllers"
)

func init() {
	// 首页
	beego.Router("/", &controllers.IndexController{})

	// 聚会主页
	beego.Router("/party/:urlCode", &controllers.PartyController{})
	// 获取投票情况
	beego.Router("/party/:urlCode/party-members", &controllers.PartyController{}, "get:PartyMembers")
}
