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

func GetUser(userId int) (model User, err error) {
	o := orm.NewOrm()
	model = User{Id: userId}
	err = o.Read(&model, "user_id")
	return
}
