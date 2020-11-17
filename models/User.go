package models

import (
	"fmt"

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

func GetUser() (User) {
	o := orm.NewOrm()
	user := User{Id: 1}

	err := o.Read(&user, "Id")
	if err != nil {
		fmt.Println("没有找到")
	}

	return user
}
