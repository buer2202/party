package models

import (
	"strings"
	"time"

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

// AddPartyMember 添加参与记录
func (p *PartyMember) AddPartyMember(partyId int, memberId int, joinPeopleNum int, canJoinDate string) bool {
	canJoinDateArr := strings.Split(canJoinDate, ",")
	qs := orm.NewOrm().QueryTable(p)

	// 先删除以前的
	_, err := qs.Filter("party_id", partyId).Filter("member_id", memberId).Delete()
	if err != nil {
		common.MyLog("db_err", err.Error())
		return false
	}

	// 准备需要的数据
	now := time.Now().Format("2006-01-02 15:04:05")
	dateNum := len(canJoinDateArr)

	// 构造数据
	var insertData = make([]PartyMember, dateNum, dateNum)
	for k, v := range canJoinDateArr {
		insertData[k] = PartyMember{
			PartyId:       partyId,
			MemberId:      memberId,
			JoinPeopleNum: joinPeopleNum,
			CanJoinDate:   v,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
	}

	// 插入数据
	orm.NewOrm().InsertMulti(dateNum, insertData)

	return true
}
