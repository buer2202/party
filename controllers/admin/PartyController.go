package admin

import (
	"github.com/astaxie/beego"
)

type PartyController struct {
	beego.Controller
}

func (c *PartyController) Create() {
    c.Layout = "layout.html"
	c.TplName = "admin/party/create.tpl"
}

func (c *PartyController) Store() {

}
