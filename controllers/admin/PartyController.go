package admin

import (
	"github.com/astaxie/beego"
)

type PartyController struct {
	beego.Controller
}

func (c *UserController) Create() {
    c.Layout = "layout.html"
	c.TplName = "user/index.tpl"
}

func (c *UserController) Store() {

}
