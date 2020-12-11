package controllers

import (
	"github.com/astaxie/beego"
	"party2202.com/common"
)

// AjaxAuthController 首页
type AjaxAuthController struct {
	beego.Controller
}

// 登录失效时，重定向到此方法
func (c *AjaxAuthController) Get() {
	c.Data["json"] = common.Ajax(0, "302跳转", beego.URLFor("AuthController.LoginForm"))
	c.ServeJSON()
}
