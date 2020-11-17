package controllers

import (
	"github.com/astaxie/beego"

	"party2202.com/models"
)

type PartyController struct {
	beego.Controller
}

func (this *PartyController) Get() {
	user := models.GetUser()
	party, err := models.GetByUrlCode(this.Ctx.Input.Param(":id"))
	if (err != nil) {
		this.Abort("404")
	}

	this.Data["user"] = user.Account
	this.Data["party"] = party
	this.TplName = "party.tpl"
}
