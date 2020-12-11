package admin

import (
    "github.com/astaxie/beego"
    "party2202.com/models"
)

type MemberController struct {
	beego.Controller
}

// 成员管理主页
func (c *MemberController) Index() {
    dataList := (&models.Member{}).GetByUserId(c.GetSession("authUser").(models.User).Id)

    c.Data["dataList"] = dataList
	c.Layout = "layout.html"
	c.TplName = "admin/member/index.tpl"
}

// 添加
func (c *MemberController) Store() {

}

// 更新
func (c *MemberController) Update() {

}

// 删除
func (c *MemberController) Destroy() {

}
