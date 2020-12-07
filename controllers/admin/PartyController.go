package admin

import (
	"github.com/astaxie/beego"
)

type PartyController struct {
	beego.Controller
}

// func (c *UserController) Create() {
// 	c.authUser = c.GetSession("authUser")
// 	if c.authUser == nil {
// 		c.Redirect(beego.URLFor("AuthController.LoginForm"), 302)
// 		c.StopRun()
// 	}
// }

// func (c *UserController) Store() {
// 	c.authUser = c.GetSession("authUser")
// 	if c.authUser == nil {
// 		c.Redirect(beego.URLFor("AuthController.LoginForm"), 302)
// 		c.StopRun()
// 	}
// }
