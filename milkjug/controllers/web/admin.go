package controllers

import "milkshake/controllers"

type WebAdminController struct {
	controllers.WebBaseController
}

func (c *WebAdminController) Index() {
	if c.IsLogin {
		c.TplName = "admin.html"
		c.Data["eventPage"] = "/web/event"
		c.Data["homePage"] = "/web/home"
	} else {
		c.Ctx.Redirect(302, "/web/login")
	}
}
