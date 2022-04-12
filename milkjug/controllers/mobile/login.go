package controllers

import (
	"encoding/json"
	"log"
	"milkshake/jwt"
	"milkshake/models"
	"net/http"
	"strconv"

	"milkshake/redis"

	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type MobileLoginController struct {
	beego.Controller
}

func (c *MobileLoginController) loginFail() {
	c.Ctx.Output.SetStatus(401)
	c.Ctx.Output.Body([]byte("Invalid UserID or Password"))
}

func (c *MobileLoginController) ServerError(err error) {
	c.Ctx.Output.Header("SetStatus", strconv.Itoa(http.StatusInternalServerError))
	log.Println(err)
	http.Error(c.Ctx.ResponseWriter, "Server is Wrong", http.StatusInternalServerError)
}

func (c *MobileLoginController) Login() {
	if c.Ctx.Input.IsPost() {
		var userLogin models.LoginTemp
		if userLogin.SensitizeLogin(c.Ctx.Input.RequestBody) == nil {
			if userLogin.Id != 0 {
				if record, err := models.GetUsersById(userLogin.Id); err == nil {
					if bcrypt.CompareHashAndPassword([]byte(record.UserPassword), []byte(userLogin.UserPassword)) == nil {
						claims := jwt.GenerateTokenMapClaims(0)
						if tokenString, err := jwt.GenerateTokenString(claims); err == nil {
							//json serialize
							var deviceInfo models.DeviceInfo
							deviceInfo.ConstructFromHeader(c.Ctx.Request.Header)
							creden := models.ApiCredentials{Id: userLogin.Id, Type: record.UserType, Device: deviceInfo}
							if by, err := json.Marshal(creden); err == nil {
								redis.SetStr(tokenString, string(by), 0)
								log.Printf("UserID %d login", userLogin.Id)
								log.Println(deviceInfo)

								c.Ctx.Output.Header("Authorization", tokenString)
								c.Ctx.Output.SetStatus(200)
								c.ServeJSON()
								return
							}
						}
						c.ServerError(err)
					} else {
						log.Println("Password is invalid")
						c.loginFail()
					}
				} else {
					log.Println("Couldn't find user with provied ID")
					c.loginFail()
				}
			} else {
				log.Println("UserID can't be 0")
				c.loginFail()
			}
		} else {
			log.Println("Unable to unpack json data")
			c.loginFail()
		}
		c.ServeJSON()
	}
}
