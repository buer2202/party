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

// 活动管理主页
func (c *PartyController) Index() {
	c.Layout = "layout.html"
	c.TplName = "admin/party/index.tpl"
}

// 翻页数据
func (c *PartyController) DataList() {
	page, _ := c.GetInt64("p", 1)
	dataList := (&models.Party{}).GetList(c.GetSession("authUser").(models.User).Id, page)

	c.Data["json"] = dataList
	c.ServeJSON()
}

// 创建页
func (c *PartyController) Create() {
	c.Layout = "layout.html"
	c.TplName = "admin/party/create.tpl"
}

// 添加活动方法
func (c *PartyController) Store() {
	valid := validation.Validation{}
	valid.Required(c.GetString("name"), "name")
	valid.Required(c.GetString("party_desc"), "party_desc")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "请填写完整", nil)
		c.ServeJSON()
		c.StopRun()
	}

	_, err := (&models.Party{}).Store(c.GetSession("authUser").(models.User).Id, c.GetString("name"), c.GetString("party_desc"))
	if err != nil {
		c.Data["json"] = common.Ajax(0, "活动创建失败", nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.Data["json"] = common.Ajax(1, "操作成功", nil)
	c.ServeJSON()
}

// 活动确认
func (c *PartyController) Confirm() {
	partyId, _ := c.GetInt("id")
	confirmDesc := c.GetString("confirm_desc")

	valid := validation.Validation{}
	valid.Required(partyId, "id")
	valid.Required(confirmDesc, "confirm_desc")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "参数错误", nil)
		c.ServeJSON()
		c.StopRun()
	}

	err := (&models.Party{}).Confirm(c.GetSession("authUser").(models.User).Id, partyId, confirmDesc)
	if err != nil {
		c.Data["json"] = common.Ajax(0, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}

	c.Data["json"] = common.Ajax(1, "操作成功", nil)
	c.ServeJSON()
}

// 分享活动链接
func (c *PartyController) ShareUrl() {
	valid := validation.Validation{}
	valid.Required(c.GetString("url_code"), "url_code")
	if valid.HasErrors() {
		c.Data["json"] = common.Ajax(0, "参数错误", nil)
		c.ServeJSON()
		c.StopRun()
	}

	uri := beego.URLFor("home.PartyController.Get", ":urlCode", c.GetString("url_code"))
	shareUrl := common.GetUrl(c.Ctx.Request, uri)
	c.Data["json"] = common.Ajax(1, "操作成功", shareUrl)
	c.ServeJSON()
}
