package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"party2202.com/models"
)

// FilterGuest 已登录跳转到用户首页
func FilterGuest(ctx *context.Context) {
	_, ok := ctx.Input.Session("authUser").(models.User)
	if ok {
		ctx.Redirect(302, beego.URLFor("admin.UserController.Get"))
	}
}

// FilterUser 未登录跳转到登录页面
func FilterUser(ctx *context.Context) {
	_, ok := ctx.Input.Session("authUser").(models.User)
	if !ok {
		ctx.Redirect(302, beego.URLFor("home.AuthController.LoginForm"))
	}
}
