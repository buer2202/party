package home

import (
	"github.com/astaxie/beego"
	"party2202.com/common"
	"party2202.com/models"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) LoginForm() {
	c.Layout = "layout.html"
	c.TplName = "home/auth/login.tpl"
}

func (c *AuthController) Login() {
	user, err := (&models.User{}).Login(c.GetString("account"), c.GetString("password"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), "")
		c.ServeJSON()
		c.StopRun()
	}

	c.SetSession("authUser", user)

	c.Data["json"] = common.Ajax(1, "登录成功", "")
	c.ServeJSON()
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Data["json"] = common.Ajax(1, "注销成功", "")
	c.Redirect(beego.URLFor("home.AuthController.LoginForm"), 302)
}
