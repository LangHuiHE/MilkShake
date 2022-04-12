package controllers

import (
	controller "milkshake/controllers"
	"milkshake/models"
)

type MobileUserController struct {
	controller.BaseController
}

func (c *MobileUserController) Get() {
	if c.Verified {
		var generalUserInfo models.GeneralUserInfo
		if err := generalUserInfo.CreateGeneralUserInfo(c.Id); err == nil {
			c.Data["json"] = generalUserInfo
			// log.Println(generalUserInfo)
		} else {
			c.ServerError(err)
		}
		if err := c.RenewToken(); err != nil {
			c.ServerError(err)
		}
	} else {
		c.Unauthorized()
	}
	c.ServeJSON()
}
