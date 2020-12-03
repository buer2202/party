package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"party2202.com/common"
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
	if err != nil {
		return data, errors.New("账号不存在")
	}

	if data.Password != common.HashSha256(password) {
		return data, errors.New("密码错误")
	}

	return
}
