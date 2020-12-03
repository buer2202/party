package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// TestController 首页
type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	fmt.Println(c.GetSession("authUser"))
	c.TplName = "test.tpl"
}
