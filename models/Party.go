package models

import (
	"time"

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

func (m *Party) GetList(userId int, pageNo int64) common.Page {
	// 分页参数
	var pageSize int64
	pageSize = 5
	offset := (pageNo - 1) * pageSize

	qs := orm.NewOrm().QueryTable(m)
	count, err0 := qs.Filter("user_id", userId).Limit(pageSize, offset).Count()
	if err0 != nil {
		common.MyLog("db_err", err0.Error())
		return common.Page{}
	}

	var dataList []*Party
	_, err1 := qs.Filter("user_id", userId).Limit(pageSize, offset).All(&dataList)
	if err1 != nil {
		common.MyLog("db_err", err1.Error())
		return common.Page{}
	}

	return common.PageUtil(count, pageNo, pageSize, dataList)
}

// GetByUrlCode 用urlCode查询
func (m *Party) GetByUrlCode(urlCode string) (data Party, err error) {
	qs := orm.NewOrm().QueryTable(m)
	err = qs.Filter("url_code", urlCode).One(&data)
	if err != nil {
		common.MyLog("db_err", err.Error())
	}

	return
}

// Create 添加参与记录
func (m *Party) Store(userId int, name string, partyDesc string) (id int64, err error) {
	// 构造数据
	now := time.Now().Format("2006-01-02 15:04:05")

	var model Party
	model.UserId = userId
	model.UrlCode = common.RandSeq(6)
	model.Name = name
	model.PartyDesc = partyDesc
	model.CreatedAt = now
	model.UpdatedAt = now

	id, err = orm.NewOrm().Insert(&model)
	return
}
