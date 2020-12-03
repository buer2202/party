package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
	authUser interface{}
}

func (c *UserController) Prepare() {
	c.authUser = c.GetSession("authUser")
	if c.authUser == nil {
		c.Redirect(beego.URLFor("AuthController.LoginForm"), 302)
		c.StopRun()
	}
}

func (c *UserController) Get() {
	c.Data["authUser"] = c.GetSession("authUser")

	c.Layout = "layout.html"
	c.TplName = "user/index.tpl"
}
