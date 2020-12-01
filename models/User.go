package models

import (
	"github.com/astaxie/beego/orm"
)

// User Model
type User struct {
	Id        int
	Account   string
	Password  string
	Nickname  string
	CreatedAt string
	UpdatedAt string
}

func (m *User) Login(account string, password string) (data User, err error) {
	qs := orm.NewOrm().QueryTable(m)
	err = qs.Filter("account", account).One(&data)

	return
}
