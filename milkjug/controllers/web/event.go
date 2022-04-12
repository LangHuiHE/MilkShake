package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type EventController struct {
	beego.Controller
}

func (c *EventController) Index() {
	c.Data["Website"] = "github.com/beego"
	c.Data["Email"] = "langhuihe@gmail.com"
	c.Data["ProjectName"] = "Milkshake"
	c.Data["Author"] = "LangHui He"
	c.Data["Github"] = "github.com/LangHuiHE"
	c.Data["LoginPage"] = "/web/login"
	c.Data["homePage"] = "/web/home"
	c.TplName = "event.html"
}
