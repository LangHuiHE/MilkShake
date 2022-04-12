package controllers

import (
	"encoding/json"
	"log"
	"milkshake/controllers"
	"milkshake/jwt"
	"milkshake/models"
	"milkshake/redis"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type WebLoginController struct {
	controllers.WebBaseController
}

func (c *WebLoginController) addQrToken() (string, string, error) {
	var empty map[string]interface{}
	tokenString, err := jwt.LazyTokenString(empty, 5)

	if err == nil {
		content := []string{c.URLFor("WebLoginController.LoginWithQrCode"), tokenString}
		qr := models.QrCode{Msg: strings.Join(content, "%20")}
		if imageString, err := qr.GeneratePNGonBase64(); err == nil {
			c.Data["QRcode"] = imageString
			log.Println(tokenString)
			redis.SetStr(tokenString, "", time.Minute*time.Duration(5))
			return tokenString, imageString, nil
		}
	}
	return "", "", err
}

func (c *WebLoginController) Index() {
	c.AddBasicInfo()
	c.TplName = "login.html"
	c.Data["eventPage"] = "/web/event"
	c.Data["homePage"] = "/web/home"
}

func (c *WebLoginController) Login() {
	if c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("WebAdminController.Index"))
		return
	}

	c.Index()

	if c.Ctx.Input.IsPost() {
		var userLogin models.LoginTemp
		if err := userLogin.SensitizeLogin(c.Ctx.Input.RequestBody); err == nil {
			if userLogin.Id != 0 {
				if record, err := models.GetUsersById(userLogin.Id); err == nil {
					if bcrypt.CompareHashAndPassword([]byte(record.UserPassword), []byte(userLogin.UserPassword)) == nil {
						claims := map[string]interface{}{
							"id": userLogin.Id,
						}
						if token, err := jwt.LazyTokenString(claims, 0); err == nil {
							log.Printf("UserID %d login on Website", userLogin.Id)
							c.IsLogin = true
							c.User = record
							if by, err := json.Marshal(record); err == nil {
								redis.SetStr(token, string(by), 0)
							} else {
								log.Println("Can't incode user to cache: ", err)
								c.ServerError(err)
							}
							c.Ctx.SetCookie("jwt", token)
							c.Ctx.Redirect(200, c.URLFor("WebAdminController.Index"))
							log.Println(c.URLFor("WebAdminController.Index"))
							//c.Ctx.Output.SetStatus(200)
							c.ServeJSON()
							return
						} else {
							log.Println("Can't create jwt", err)
							c.ServerError(err)
						}
					} else {
						log.Println("Password is invalid: ", err)
					}
				} else {
					log.Println("Couldn't find user with provied ID: ", err)
				}
			} else {
				log.Println("UserID can't be 0")

			}
		} else {
			log.Println("Unable to unpack json data: ", err)
		}
		c.LoginFail()
		c.ServeJSON()
	}
}

func (c *WebLoginController) LoginWithQrCode() {
	if c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("WebAdminController.Index"))
		return
	}

	if c.Ctx.Input.IsGet() {
		if qrToken := c.Ctx.Input.Header("QrToken"); qrToken == "" {
			// get the Qr Code
			if tokenString, imageString, err := c.addQrToken(); err == nil {
				c.Ctx.Output.SetStatus(200)
				c.Ctx.Output.Header("QrToken", tokenString)
				c.Data["json"] = imageString
				c.ServeJSON()
				return
			} else {
				c.ServerError(err)
				return
			}
		} else {
			// has QrToken
			// check status
			// log.Println("token from api call ", qrToken)
			// log.Println(redis.GetStr(qrToken))
			if v := redis.GetStr(qrToken); v == "" {
				// hasn't verified user
			} else if statusValid, _ := jwt.CheckToken(v); statusValid {
				c.Ctx.SetCookie("jwt", v)
				c.IsLogin = true
				c.Ctx.Redirect(200, c.URLFor("WebAdminController.Index"))
				// log.Println(c.URLFor("WebAdminController.Index"))
				return
			}
		}
	} else if c.Ctx.Input.IsPost() {
		if token, qrToken := c.Ctx.Input.Header("Authorization"), c.Ctx.Input.Header("QrToken"); token != "" && qrToken != "" {
			if tokenOk, _ := jwt.CheckToken(token); tokenOk {
				if qrTokenOk, _ := jwt.CheckToken(qrToken); qrTokenOk {
					if record := redis.GetByte(token); len(record) > 0 {
						var recordCreden models.ApiCredentials
						json.Unmarshal(record, &recordCreden)
						var requestDevice models.DeviceInfo
						requestDevice.ConstructFromHeader(c.Ctx.Request.Header)
						if recordCreden.Device == requestDevice && redis.GetStr(qrToken) == "" {
							userID := recordCreden.Id
							// userType := recordCreden.Type

							log.Printf("UserID %d verified by QrToken", userID)
							claims := map[string]interface{}{
								"id": userID,
							}
							if webSessionToken, err := jwt.LazyTokenString(claims, 0); err == nil {
								log.Println("crated jwt for user: ", userID)
								redis.SetStr(qrToken, webSessionToken, time.Minute*time.Duration(2))
								// log.Println(redis.GetStr(qrToken))

								// Renew Auth Token
								new, err := jwt.RenewToken(token, 0, time.Duration(2*time.Second))
								if err == nil {
									c.Ctx.Output.Header("Authorization", new)
									c.Ctx.Output.SetStatus(200)
									return
								} else {
									c.ServerError(err)
									return
								}
							} else {
								log.Println("Can't create jwt", err)
								c.ServerError(err)
							}
						}
					} else {
						log.Println("LoginWithQrCode() can't find cache record")
					}
				} else {
					log.Println("LoginWithQrCode() QrToken fail")
				}
			} else {
				log.Println("LoginWithQrCode() Authorization fail")
			}
		} else {
			log.Println("LoginWithQrCode() Can't find tokens in header")
			log.Println("auth token:", token)
			log.Println("qr token:", qrToken)
		}
	}
	c.LoginFail()
}

func (c *WebLoginController) Logout() {
	c.Ctx.Redirect(302, c.URLFor("WebLoginController.Login"))
}
