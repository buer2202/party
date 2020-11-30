package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"

	"party2202.com/models"
)

// PartyController 聚会控制器
type PartyController struct {
	beego.Controller
}

// Get 主页
func (p *PartyController) Get() {
	party, err := (&models.Party{}).GetByUrlCode(p.Ctx.Input.Param(":urlCode"))
	if err != nil {
		p.Abort("404")
	}
	members := models.MemberGetByUserId(party.UserId)
	partyMemberDate := (&models.PartyMember{}).GetPartyDate(party.Id)

	p.Data["urlCode"] = p.Ctx.Input.Param(":urlCode")
	p.Data["party"] = party
	p.Data["members"] = members
	p.Data["partyMemberDate"] = partyMemberDate
	p.TplName = "party.tpl"
}

// PartyMembers 获取活动成员
func (p *PartyController) PartyMembers() {
	party, err := (&models.Party{}).GetByUrlCode(p.Ctx.Input.Param(":urlCode"))
	if err != nil {
		p.Data["json"] = [0]int{}
	} else {
		p.Data["json"] = (&models.PartyMember{}).GetPartyMember(party.Id)
	}

	p.ServeJSON()
}

// Post 参与提交
func (p *PartyController) Post() {
	memberId, _ := p.GetInt("member_id")
	joinPeopleNum, _ := p.GetInt("join_people_num")
	canJoinDate := p.GetString("can_join_date")

	valid := validation.Validation{}
	valid.Required(memberId, "member_id")
	valid.Required(joinPeopleNum, "join_people_num")
	valid.Required(canJoinDate, "can_join_date")
	if valid.HasErrors() {
		p.Data["json"] = false
		p.ServeJSON()
		// for _, err := range valid.Errors {
		// 	log.Println(err.Key, err.Message)
		// }
	}

	party, err := (&models.Party{}).GetByUrlCode(p.Ctx.Input.Param(":urlCode"))
	if err != nil {
		p.Data["json"] = false
		p.ServeJSON()
	}

	p.Data["json"] = (&models.PartyMember{}).AddPartyMember(
		party.Id,
		memberId,
		joinPeopleNum,
		canJoinDate,
	)

	p.ServeJSON()
}
