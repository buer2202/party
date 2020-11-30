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

// GetPartyDate 根据PartyID获取所有报名的日期，不重复
func (p *PartyMember) GetPartyDate(partyId int) (dataList []*PartyMember) {
	qs := orm.NewOrm().QueryTable(p)
	_, err := qs.Distinct().Filter("party_id", partyId).OrderBy("can_join_date").All(&dataList, "can_join_date")
	if err != nil {
		common.MyLog("db_err", err.Error())
	}
	return
}

// GetPartyMember 根据PartyID获取参与人员
func (p *PartyMember) GetPartyMember(partyId int) (dataList []*PartyMember) {
	qs := orm.NewOrm().QueryTable(p)
	_, err := qs.Filter("party_id", partyId).All(&dataList, "party_id", "member_id", "join_people_num", "can_join_date")
	if err != nil {
		common.MyLog("db_err", err.Error())
	}
	return
}
