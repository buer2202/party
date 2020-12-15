package models

import (
	"errors"
	"fmt"
	"time"

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

func (m *User) Regist(account string, password string, nickname string) (data User, err error) {
	qs := orm.NewOrm().QueryTable(m)
	err = qs.Filter("account", account).One(&data)
	fmt.Println(data)
	if err == nil {
		return data, errors.New("账号已存在")
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	var model User
	model.Account = account
	model.Password = common.HashSha256(password)
	model.Nickname = nickname
	model.CreatedAt = now
	model.UpdatedAt = now

	_, err = orm.NewOrm().Insert(&model)
	if err != nil {
		return model, errors.New("注册失败")
	}

	return model, nil
}
