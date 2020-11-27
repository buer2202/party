package models

import (
	"github.com/astaxie/beego/orm"
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

// PartyGetByUrlCode 用urlCode查询
func PartyGetByUrlCode(urlCode string) (model Party, err error) {
	o := orm.NewOrm()
	model = Party{UrlCode: urlCode}
	err = o.Read(&model, "url_code")
	return
}
