package admin

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"party2202.com/common"
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
	valid := validation.Validation{}
	valid.Required(c.GetString("nickname"), "nickname")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "请填写昵称", nil)
		c.ServeJSON()
		c.StopRun()
	}

	_, err := (&models.Member{}).Create(c.GetSession("authUser").(models.User).Id, c.GetString("nickname"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, "成员创建失败", nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.Data["json"] = common.Ajax(1, "操作成功", nil)
	c.ServeJSON()
}

// 更新
func (c *MemberController) Update() {
	valid := validation.Validation{}
	valid.Required(c.GetString("nickname"), "nickname")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "请填写昵称", nil)
		c.ServeJSON()
		c.StopRun()
	}

	memberId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}

	err = (&models.Member{}).Update(c.GetSession("authUser").(models.User).Id, memberId, c.GetString("nickname"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.Data["json"] = common.Ajax(1, "操作成功", nil)
	c.ServeJSON()
}

// 删除
func (c *MemberController) Destroy() {
	memberId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}

	err = (&models.Member{}).Delete(c.GetSession("authUser").(models.User).Id, memberId)
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.Data["json"] = common.Ajax(1, "操作成功", nil)
	c.ServeJSON()
}
