package admin

import (
	"github.com/astaxie/beego"
)

type MemberController struct {
	beego.Controller
}

// 成员管理主页
func (c *MemberController) Index() {
	c.Layout = "layout.html"
	c.TplName = "admin/member/index.tpl"
}
