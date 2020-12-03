package routers

import (
	"github.com/astaxie/beego"
	"party2202.com/controllers"
)

func init() {
	// 错误页面
	beego.ErrorController(&controllers.ErrorController{})

	// 欢迎页
	beego.Router("/", &controllers.IndexController{})

	// 聚会主页
	beego.Router("/party/:urlCode", &controllers.PartyController{})
	// 获取参与情况
	beego.Router("/party/:urlCode/party-members", &controllers.PartyController{}, "get:PartyMembers")
	// 提交参与
	beego.Router("/party/:urlCode", &controllers.PartyController{})

	// 用户认证
	beego.Router("/login", &controllers.AuthController{}, "get:LoginForm")
	beego.Router("/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/logout", &controllers.AuthController{}, "post:Logout")

	// 用户前台
	beego.Router("/user", &controllers.UserController{})
}
