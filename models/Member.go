package models

import (
	"github.com/astaxie/beego/orm"
)

// Member Model
type Member struct {
	Id        int
	UserId    int
	Nickname  string
	CreatedAt string
	UpdatedAt string
}

func GetByUserId(userId int) []*Member {
	o := orm.NewOrm()
    model := new(Member)
    var dataList []*Member
	o.QueryTable(model).Filter("user_id", userId).All(dataList)
	return dataList
}
