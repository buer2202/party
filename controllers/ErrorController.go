package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.Layout = "layout.html"
	c.TplName = "error/404.tpl"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.Layout = "layout.html"
	c.TplName = "error/501.tpl"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.Layout = "layout.html"
	c.TplName = "error/db.tpl"
}
