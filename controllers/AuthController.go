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
	_, err := (&models.User{}).Login(c.GetString("account"), c.GetString("password"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, "数据写入失败", "")
		c.ServeJSON()
    }

	c.Data["json"] = common.Ajax(1, "提交成功", "")
	c.ServeJSON()
}

func (c *AuthController) Logout() {

}
