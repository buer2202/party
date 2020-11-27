package models

import (
	"github.com/astaxie/beego/orm"
	"party2202.com/common"
)

type Member struct {
	Id        int
	UserId    int
	Nickname  string
	CreatedAt string
	UpdatedAt string
}

func MemberGetByUserId(userId int) (dataList []*Member) {
	o := orm.NewOrm()
	model := new(Member)
	_, err := o.QueryTable(model).Filter("user_id", userId).All(&dataList)
	if err != nil {
		common.MyLog("db_err", err.Error())
	}

	return
}
