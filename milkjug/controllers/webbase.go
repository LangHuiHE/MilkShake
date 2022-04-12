package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"milkshake/jwt"
	"milkshake/models"
	"milkshake/redis"
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type WebBaseController struct {
	beego.Controller
	IsLogin bool
	Token   string
	User    *models.Users
	QrToken *models.QrCode
}

func (c *WebBaseController) Prepare() {
	c.checkIsLogin()
}

func (c *WebBaseController) checkIsLogin() {
	if token := c.Ctx.Input.Cookie("jwt"); token != "" {
		// log.Println(token)
		if ok, _ := jwt.CheckToken(token); ok {
			if record := redis.GetByte(token); len(record) > 0 {
				var user models.Users
				json.Unmarshal(record, &user)
				c.User = &user
				c.IsLogin = true

				log.Println("User: ", c.User.Id, "already login")
				return
			} else {
				// token valid but can't record
				// store record into cache
				if _, claims, ok := jwt.ReadToken(token); ok {
					for k, v := range *claims {
						if k == "id" {
							var id int
							switch v.(type) {
							case string:
								s := v.(string)
								id, _ = strconv.Atoi(s)
							case int:
								id = v.(int)
							case float64:
								id = int(v.(float64))
							default:
								fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, v)
								break
							}

							if user, err := models.GetUsersById(id); err == nil {
								if by, err := json.Marshal(user); err == nil {
									redis.SetStr(token, string(by), 0)
									c.User = user
									c.IsLogin = true

									log.Println("User: ", c.User.Id, "already login")
									return
								} else {
									log.Println("Can't convert incode user to cache: ", err)
								}
							} else {
								log.Println("Can't get User: ", err)
							}
						}
					}
					log.Println("Can't find (id) in jwt")
				} else {
					log.Println("Can't decode jwt")
				}
			}
		}
	}
	// c.IsLogin = false
}

func (c *WebBaseController) ServerError(err error) {
	c.Ctx.Output.Header("SetStatus", strconv.Itoa(http.StatusInternalServerError))
	log.Println(err)
	http.Error(c.Ctx.ResponseWriter, "Server is Wrong", http.StatusInternalServerError)
}

func (c *WebBaseController) Unauthorized() {
	c.Ctx.Output.SetStatus(401)
	c.Ctx.Output.Body([]byte("Unauthorized Request"))
}

func (c *WebBaseController) LoginFail() {
	c.Ctx.Output.SetStatus(401)
	c.Ctx.Output.Body([]byte("Invalid UserID or Password"))
}

func (c *WebBaseController) AddBasicInfo() {
	c.Data["Website"] = "github.com/beego"
	c.Data["Email"] = "langhuihe@gmail.com"
	c.Data["ProjectName"] = "Milkshake"
	c.Data["Author"] = "LangHui He"
	c.Data["Github"] = "github.com/LangHuiHE"
}
