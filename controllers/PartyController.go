package controllers

import (
	"github.com/astaxie/beego"

	"party2202.com/models"
)

// PartyController 聚会控制器
type PartyController struct {
	beego.Controller
}

func (this *PartyController) Get() {
	party, err := models.GetByUrlCode(this.Ctx.Input.Param(":id"))
	if err != nil {
		this.Abort("404")
	}
	members := models.GetByUserId(party.UserId)

	this.Data["party"] = party
	this.Data["members"] = members
	this.TplName = "party.tpl"
}
