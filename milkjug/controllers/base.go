package controllers

import (
	"encoding/json"
	"log"
	"milkshake/jwt"
	"milkshake/models"
	"milkshake/redis"
	"net/http"
	"strconv"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	Verified bool
	Token    string
	Id       int
	Type     int
}

func (c *BaseController) Prepare() {
	c.Verified = c.VerifyApiCredentials()
}

func (c *BaseController) ServerError(err error) {
	c.Ctx.Output.Header("SetStatus", strconv.Itoa(http.StatusInternalServerError))
	log.Println(err)
	http.Error(c.Ctx.ResponseWriter, "Server is Wrong", http.StatusInternalServerError)
}

func (c *BaseController) VerifyApiCredentials() bool {
	if token := c.Ctx.Input.Header("Authorization"); token != "" {
		if ok, _ := jwt.CheckToken(token); ok {
			if record := redis.GetByte(token); len(record) > 0 {
				var recordCreden models.ApiCredentials
				json.Unmarshal(record, &recordCreden)
				var requestDevice models.DeviceInfo
				requestDevice.ConstructFromHeader(c.Ctx.Request.Header)
				// log.Println("requestDevice ", requestDevice)
				// log.Println("record ", recordDevice)
				if recordCreden.Device == requestDevice {
					c.Token = token
					c.Id = recordCreden.Id
					c.Type = recordCreden.Type
					log.Printf("UserID %d verified", c.Id)
					return true
				}

			}
			log.Println("checkToken is ", ok)
			log.Println("Get Record", redis.GetByte(token))
			// log.Println("Swapkey.GetStr Newkey:", redis.GetStr(token))
		}
	}
	log.Println(c.Ctx.Input.Header("Authorization"))
	return false
}

func (c *BaseController) Unauthorized() {
	c.Ctx.Output.SetStatus(401)
	c.Ctx.Output.Body([]byte("Unauthorized Request"))
}

func (c *BaseController) RenewToken() error {
	new, err := jwt.RenewToken(c.Token, 0, time.Duration(2*time.Second))
	if err == nil {
		c.Ctx.Output.Header("Authorization", new)
		c.Ctx.Output.SetStatus(200)
		return nil
	}
	return err
}
