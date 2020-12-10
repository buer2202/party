package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"party2202.com/common"
	"party2202.com/models"
)

type PartyController struct {
	beego.Controller
}

func (c *PartyController) Index() {
	c.Layout = "layout.html"
	c.TplName = "admin/party/index.tpl"
}

func (c *PartyController) DataList() {
	page, _ := c.GetInt64("p", 1)
	dataList := (&models.Party{}).GetList(c.GetSession("authUser").(models.User).Id, page)

	c.Data["json"] = dataList
	c.ServeJSON()
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
		c.Data["json"] = common.Ajax(0, "请填写完整", nil)
		c.ServeJSON()
		c.StopRun()
	}

	userInfo := c.GetSession("authUser")
	_, err := (&models.Party{}).Store(userInfo.(models.User).Id, c.GetString("name"), c.GetString("party_desc"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, "活动创建失败", nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.Data["json"] = common.Ajax(1, "操作成功", nil)
	c.ServeJSON()
}

func (c *PartyController) Update() {
	c.Layout = "layout.html"
	c.TplName = "admin/party/create.tpl"
}
