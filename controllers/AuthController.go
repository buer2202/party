package controllers

import (
	"github.com/astaxie/beego"
	"party2202.com/common"
	"party2202.com/models"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) LoginForm() {
	c.TplName = "auth/login.tpl"
}

func (c *AuthController) Login() {
	user, err := (&models.User{}).Login(c.GetString("account"), c.GetString("password"))
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		c.StopRun()
	}

	c.SetSession("authUser", user)

	c.Data["json"] = common.Ajax(1, "登录成功", "")
	c.ServeJSON()
}

func (c *AuthController) Logout() {
	c.DestroySession()
}
