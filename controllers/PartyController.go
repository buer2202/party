package controllers

import (
	"github.com/astaxie/beego"
)

type PartyController struct {
	beego.Controller
}

func (c *PartyController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "party.tpl"
}
