package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HomePageController struct {
	beego.Controller
}

func (c *HomePageController) Index() {
	c.Data["Website"] = "github.com/beego"
	c.Data["Email"] = "langhuihe@gmail.com"
	c.Data["ProjectName"] = "Milkshake"
	c.Data["Author"] = "LangHui He"
	c.Data["Github"] = "github.com/LangHuiHE"
	// IDK why can;t generate url maybe one is base controller one is not?
	// log.Println(beego.URLFor("WebLoginController.Index"))
	c.Data["LoginPage"] = "/web/login"
	c.Data["EventPage"] = "/web/event"
	c.Data["Info"] = "There are five applications related with Dixie State University or Utah Tech University in IOS platform which is App Store, Trailblazer Nation, Dixie State Traditions, Trailblazer Athletics, DSU Safe, Dixie Rec. All of them have not good ratings and only have few ratings reports. Thus, those applications or the information adn function they provide are not attracted to students. To sum up, my proposal application is having the most important functionalities of those five apps in the App Store with other new features I state below and will become more popular with the university's students."
	c.TplName = "homePage.html"
}
