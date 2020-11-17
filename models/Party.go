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

func GetByUrlCode(urlCode string) (model Party, err error) {
	o := orm.NewOrm()
	model = Party{UrlCode: urlCode}
	err = o.Read(&model, "urlCode")
	return
}
