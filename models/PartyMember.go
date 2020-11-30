package models

import (
	"github.com/astaxie/beego/orm"
	"party2202.com/common"
)

type PartyMember struct {
	Id            int
	PartyId       int
	MemberId      int
	JoinPeopleNum int
	CanJoinDate   string
	CreatedAt     string
	UpdatedAt     string
}

// GetPartyMemberByPartyId 根据PartyID获取所有报名的日期，不重复
func PartyMemberGetDateByPartyId(partyId int) (dataList []*PartyMember) {
	o := orm.NewOrm()
	model := new(PartyMember)
	_, err := o.QueryTable(model).Distinct().Filter("party_id", partyId).OrderBy("can_join_date").All(&dataList, "can_join_date")
	if err != nil {
		common.MyLog("db_err", err.Error())
	}
	return
}

// PartyMemberGetByPartyId 根据PartyID获取参与人员
func PartyMemberGetByPartyId(partyId int) (dataList []*PartyMember) {
	o := orm.NewOrm()
	model := new(PartyMember)
	_, err := o.QueryTable(model).Filter("party_id", partyId).All(&dataList, "party_id", "member_id", "join_people_num", "can_join_date")
	if err != nil {
		common.MyLog("db_err", err.Error())
	}
	return
}
