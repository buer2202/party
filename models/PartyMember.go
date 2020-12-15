package models

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"party2202.com/common"
)

type PartyMember struct {
	Id             int
	PartyId        int
	MemberNickname string
	JoinPeopleNum  int
	CanJoinDate    string
	CreatedAt      string
	UpdatedAt      string
}

// GetPartyDate 根据PartyID获取所有报名的日期，不重复
func (m *PartyMember) GetPartyDate(partyId int) (dataList []*PartyMember) {
	qs := orm.NewOrm().QueryTable(m)
	_, err := qs.Distinct().Filter("party_id", partyId).OrderBy("can_join_date").All(&dataList, "can_join_date")
	if err != nil {
		common.MyLog("db_err", err.Error())
		return nil
	}
	return
}

// 根据partyId获取已报名的成员
func (m *PartyMember) GetJoinedMember(partyId int) (dataList []*PartyMember) {
	qs := orm.NewOrm().QueryTable(m)
	_, err := qs.Distinct().Filter("party_id", partyId).OrderBy("member_nickname").All(&dataList, "member_nickname")
	if err != nil {
		common.MyLog("db_err", err.Error())
		return nil
	}
	return
}

// GetPartyMember 根据PartyID获取参与人员
func (m *PartyMember) GetPartyMember(partyId int) (dataList []*PartyMember) {
	qs := orm.NewOrm().QueryTable(m)
	_, err := qs.Filter("party_id", partyId).All(&dataList, "party_id", "member_nickname", "join_people_num", "can_join_date")
	if err != nil {
		common.MyLog("db_err", err.Error())
		return nil
	}
	return
}

// AddPartyMember 添加参与记录
func (m *PartyMember) AddPartyMember(partyId int, memberId int, joinPeopleNum int, canJoinDate string) error {
	// 获取成员
	member, err0 := (&Member{}).GetById(memberId)
	if err0 != nil {
		return errors.New("成员不存在")
	}

	canJoinDateArr := strings.Split(canJoinDate, ",")
	qs := orm.NewOrm().QueryTable(m)

	// 先删除以前的
	_, err1 := qs.Filter("party_id", partyId).Filter("member_nickname", member.Nickname).Delete()
	if err1 != nil {
		common.MyLog("db_err", err1.Error())
		return errors.New("清除数据失败")
	}

	// 准备需要的数据
	now := time.Now().Format("2006-01-02 15:04:05")
	dateNum := len(canJoinDateArr)

	// 构造数据
	var insertData = make([]PartyMember, dateNum, dateNum)
	for k, v := range canJoinDateArr {
		insertData[k] = PartyMember{
			PartyId:        partyId,
			MemberNickname: member.Nickname,
			JoinPeopleNum:  joinPeopleNum,
			CanJoinDate:    v,
			CreatedAt:      now,
			UpdatedAt:      now,
		}
	}

	// 插入数据
	orm.NewOrm().InsertMulti(dateNum, insertData)

	return nil
}
