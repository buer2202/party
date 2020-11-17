package controllers

import (
	"github.com/astaxie/beego"
)

// IndexController 首页
type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.TplName = "index.tpl"
}
