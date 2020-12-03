package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

// IndexController 首页
type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	fmt.Println(c.GetSession("authUser"))
	c.TplName = "index.tpl"
}
