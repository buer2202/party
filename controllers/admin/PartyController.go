package admin

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"party2202.com/common"
	"party2202.com/models"
)

type PartyController struct {
	beego.Controller
}

func (c *PartyController) Create() {
	c.Layout = "layout.html"
	c.TplName = "admin/party/create.tpl"
}

func (c *PartyController) Store() {
	valid := validation.Validation{}
	valid.Required(c.GetString("name"), "name")
	valid.Required(c.GetString("party_desc"), "party_desc")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "请填写完整", "")
		c.ServeJSON()
		c.StopRun()
	}

	userInfo := c.GetSession("authUser")
	fmt.Println(userInfo)
	_, err := (&models.Party{}).Store(userInfo.(models.User).Id, c.GetString("name"), c.GetString("party_desc"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, "活动创建失败", "")
		c.ServeJSON()
		c.StopRun()
	}

	c.Data["json"] = common.Ajax(1, "操作成功", "")
	c.ServeJSON()
}
