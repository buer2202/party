package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // mysql
)

func init() {
	// orm.Debug = true
	dataSource := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass")
	dataSource += "@tcp(" + beego.AppConfig.String("mysqlurls") + ")"
	dataSource += "/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)

	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(User),
		new(Party),
		new(Member),
		new(PartyMember),
	)
}
