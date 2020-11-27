package controllers

import (
	"github.com/astaxie/beego"

	"party2202.com/models"
)

// PartyController 聚会控制器
type PartyController struct {
	beego.Controller
}

// Get 主页
func (p *PartyController) Get() {
	
	party, err := models.GetByUrlCode(p.Ctx.Input.Param(":id"))
	if err != nil {
		p.Abort("404")
	}
	members := models.GetByUserId(party.UserId)

	p.Data["party"] = party
	p.Data["members"] = members
	p.TplName = "party.tpl"
}
