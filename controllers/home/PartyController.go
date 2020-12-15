package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"

	"party2202.com/common"
	"party2202.com/models"
)

// PartyController 聚会控制器
type PartyController struct {
	beego.Controller
}

// Get 主页
func (c *PartyController) Get() {
	party, err := (&models.Party{}).GetByUrlCode(c.Ctx.Input.Param(":urlCode"))
	if err != nil {
		c.Abort("404")
	}
	allMembers := (&models.Member{}).GetByUserId(party.UserId)
	joinedMembers := (&models.PartyMember{}).GetJoinedMember(party.Id)
	partyMemberDate := (&models.PartyMember{}).GetPartyDate(party.Id)

	c.Data["urlCode"] = c.Ctx.Input.Param(":urlCode")
	c.Data["party"] = party
	c.Data["allMembers"] = allMembers
	c.Data["joinedMembers"] = joinedMembers
	c.Data["partyMemberDate"] = partyMemberDate

	c.Layout = "layout.html"
	c.TplName = "home/party/index.tpl"
}

// PartyMembers 获取活动成员
func (c *PartyController) PartyMembers() {
	party, err := (&models.Party{}).GetByUrlCode(c.Ctx.Input.Param(":urlCode"))
	if err != nil {
		c.Data["json"] = [0]int{}
	} else {
		c.Data["json"] = (&models.PartyMember{}).GetPartyMember(party.Id)
	}

	c.ServeJSON()
}

// Post 参与提交
func (c *PartyController) Post() {
	memberId, _ := c.GetInt("member_id")
	joinPeopleNum, _ := c.GetInt("join_people_num")
	canJoinDate := c.GetString("can_join_date")

	valid := validation.Validation{}
	valid.Required(memberId, "member_id")
	valid.Required(joinPeopleNum, "join_people_num")
	valid.Required(canJoinDate, "can_join_date")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "请填写完整", nil)
		c.ServeJSON()
		c.StopRun()
	}

	party, err := (&models.Party{}).GetByUrlCode(c.Ctx.Input.Param(":urlCode"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, "获取活动ID失败", nil)
		c.ServeJSON()
		c.StopRun()
	}

	if party.ConfirmDesc != "" {
		c.Data["json"] = common.Ajax(0, "该活动已停止报名", nil)
		c.ServeJSON()
		c.StopRun()
	}

	rlst := (&models.PartyMember{}).AddPartyMember(
		party.Id,
		memberId,
		joinPeopleNum,
		canJoinDate,
	)
	if rlst != nil {
		c.Data["json"] = common.Ajax(0, "数据写入失败", nil)
		c.ServeJSON()
	}

	c.Data["json"] = common.Ajax(1, "提交成功", nil)
	c.ServeJSON()
}
