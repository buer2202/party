package routers

import (
	"github.com/astaxie/beego"
	"party2202.com/controllers"
	"party2202.com/controllers/home"
	"party2202.com/controllers/admin"
)

func init() {
	// 错误页
	beego.ErrorController(&controllers.ErrorController{})
	// 测试页
	beego.Router("/test", &controllers.TestController{})

	// 聚会主页
	beego.Router("/party/:urlCode", &home.PartyController{})
	// 获取参与情况
	beego.Router("/party/:urlCode/party-members", &home.PartyController{}, "get:PartyMembers")
	// 提交参与
	beego.Router("/party/:urlCode", &home.PartyController{})

	// 用户认证
	beego.Router("/", &home.AuthController{}, "get:LoginForm")
	beego.Router("/login", &home.AuthController{}, "post:Login")
	beego.Router("/logout", &home.AuthController{}, "post:Logout")

	// 用户前台
	// 导航页
	beego.Router("/admin/user", &admin.UserController{})

	// 聚会管理
	beego.Router("/admin/party/create", &admin.PartyController{}, "get:Create")
}
