package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // mysql
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(10.0.31.254:3306)/party?charset=utf8")

	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(User),
		new(Party),
	)
}
