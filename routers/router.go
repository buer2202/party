package routers

import (
	"github.com/astaxie/beego"
	"party2202.com/controllers"
	"party2202.com/controllers/admin"
	"party2202.com/controllers/home"
)

func init() {
	// 错误页
	beego.ErrorController(&controllers.ErrorController{})
	// 测试页
	beego.Router("/test", &controllers.TestController{})
	// ajax响应
	beego.Router("/ajax-auth", &controllers.AjaxAuthController{})

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
	beego.Router("/admin/party", &admin.PartyController{}, "get:Index")
	beego.Router("/admin/party/data-list", &admin.PartyController{}, "get:DataList")
	beego.Router("/admin/party/create", &admin.PartyController{}, "get:Create")
	beego.Router("/admin/party/store", &admin.PartyController{}, "post:Store")
	beego.Router("/admin/party/confirm", &admin.PartyController{}, "post:Confirm")
	beego.Router("/admin/party/share-url", &admin.PartyController{}, "get:ShareUrl")

	// 人员管理
	beego.Router("admin/member", &admin.MemberController{}, "get:Index")
	beego.Router("admin/member/store", &admin.MemberController{}, "post:Store")
	beego.Router("admin/member/:id/update", &admin.MemberController{}, "post:Update")
	beego.Router("admin/member/:id/destroy", &admin.MemberController{}, "post:Destroy")
}
