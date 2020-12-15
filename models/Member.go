package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	"party2202.com/common"
)

type Member struct {
	Id        int
	UserId    int
	Nickname  string
	CreatedAt string
	UpdatedAt string
}

func (m *Member) GetById(id int) (data Member, err error) {
	qs := orm.NewOrm().QueryTable(m)
	err = qs.Filter("id", id).One(&data)
	return
}

func (m *Member) GetByUserId(userId int) (dataList []*Member) {
	qs := orm.NewOrm().QueryTable(m)
	_, err := qs.Filter("user_id", userId).OrderBy("nickname").All(&dataList)
	if err != nil {
		common.MyLog("db_err", err.Error())
		return nil
	}

	return
}

func (m *Member) Create(userId int, nickname string) (id int64, err error) {
	// 构造数据
	now := time.Now().Format("2006-01-02 15:04:05")

	var model Member
	model.UserId = userId
	model.Nickname = nickname
	model.CreatedAt = now
	model.UpdatedAt = now

	id, err = orm.NewOrm().Insert(&model)
	return
}

func (m *Member) Update(userId int, id int, nickname string) error {
	var model Member
	qs := orm.NewOrm().QueryTable(m)
	err1 := qs.Filter("id", id).Filter("user_id", userId).One(&model)
	if err1 != nil {
		return errors.New("该成员不存在")
	}

	num, err2 := qs.Filter("id", id).Update(orm.Params{
		"nickname":   nickname,
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	})
	if num == 0 || err2 != nil {
		return errors.New("数据更新失败")
	}

	return nil
}

func (m *Member) Delete(userId int, id int) error {
	var model Member
	qs := orm.NewOrm().QueryTable(m)
	err1 := qs.Filter("id", id).Filter("user_id", userId).One(&model)
	if err1 != nil {
		return errors.New("该成员不存在")
	}

	num, err2 := qs.Filter("id", id).Delete()
	if num == 0 || err2 != nil {
		return errors.New("数据删除失败")
	}

	return nil
}
