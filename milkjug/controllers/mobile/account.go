package controllers

import (
	controller "milkshake/controllers"
	"milkshake/models"
)

type MobileAccountController struct {
	controller.BaseController
}

func (c *MobileAccountController) Get() {
	if c.Verified {
		var a models.Accounts
		if err := a.LoadAllAccountPayment(c.Id); err == nil {
			c.Data["json"] = a
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
