package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["authUser"] = c.GetSession("authUser")
	c.Layout = "layout.html"
	c.TplName = "user/index.tpl"
}
