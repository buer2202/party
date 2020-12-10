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

func (m *Member) ByUserId(userId int) (dataList []*Member) {
	qs := orm.NewOrm().QueryTable(m)
	_, err := qs.Filter("user_id", userId).All(&dataList)
	if err != nil {
		common.MyLog("db_err", err.Error())
		return nil
	}

	return
}
