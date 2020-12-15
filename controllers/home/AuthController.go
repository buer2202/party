package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"party2202.com/common"
	"party2202.com/models"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) LoginForm() {
	c.Layout = "layout.html"
	c.TplName = "home/auth/login-form.tpl"
}

func (c *AuthController) Login() {
	valid := validation.Validation{}
	valid.Required(c.GetString("account"), "account")
	valid.Required(c.GetString("password"), "password")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "请填写完整", nil)
		c.ServeJSON()
		c.StopRun()
	}

	user, err := (&models.User{}).Login(c.GetString("account"), c.GetString("password"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.SetSession("authUser", user)

	c.Data["json"] = common.Ajax(1, "登录成功", nil)
	c.ServeJSON()
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Data["json"] = common.Ajax(1, "注销成功", nil)
	c.Redirect(beego.URLFor("home.AuthController.LoginForm"), 302)
}

func (c *AuthController) Register() {
	c.Layout = "layout.html"
	c.TplName = "home/auth/register.tpl"
}

func (c *AuthController) Regist() {
	valid := validation.Validation{}
	valid.Required(c.GetString("account"), "account")
	valid.Required(c.GetString("password"), "password")
	valid.Required(c.GetString("nickname"), "nickname")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "请填写完整", nil)
		c.ServeJSON()
		c.StopRun()
	}

	user, err := (&models.User{}).Regist(c.GetString("account"), c.GetString("password"), c.GetString("nickname"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.SetSession("authUser", user)

	c.Data["json"] = common.Ajax(1, "注册成功", nil)
	c.ServeJSON()
}
