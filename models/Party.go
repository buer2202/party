package models

import (
	"github.com/astaxie/beego/orm"
	"party2202.com/common"
)

// Party Model
type Party struct {
	Id          int
	UserId      int
	UrlCode     string
	Name        string
	PartyDesc   string
	ConfirmDesc string
	CreatedAt   string
	UpdatedAt   string
}

// GetByUrlCode 用urlCode查询
func (p *Party) GetByUrlCode(urlCode string) (data Party, err error) {
	qs := orm.NewOrm().QueryTable(p)
	err = qs.Filter("url_code", urlCode).One(&data)
	if err != nil {
		common.MyLog("db_err", err.Error())
	}

	return
}
