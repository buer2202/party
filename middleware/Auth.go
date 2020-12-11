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
		ctx.Redirect(302, beego.URLFor("UserController.Get"))
	}
}

// FilterUser 未登录跳转到登录页面
func FilterUser(ctx *context.Context) {
	_, ok := ctx.Input.Session("authUser").(models.User)
	if !ok {
		// 判断是否为ajax请求
		if ctx.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
			ctx.Redirect(302, beego.URLFor("AjaxAuthController.Get"))
		} else {
			ctx.Redirect(302, beego.URLFor("AuthController.LoginForm"))
		}
	}
}
