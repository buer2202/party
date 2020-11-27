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

func PartyMemberGetDateByPartyId(partyId int) (dataList []*PartyMember) {
	o := orm.NewOrm()
	model := new(PartyMember)
	_, err := o.QueryTable(model).Filter("party_id", partyId).All(&dataList, "party_id", "member_id", "join_people_num", "can_join_date")
	if err != nil {
		common.MyLog("db_err", err.Error())
	}
	return
}
