package controllers

import (
	"errors"
	"fmt"
	"log"
	controller "milkshake/controllers"
	"milkshake/models"
	"strconv"
	"strings"
	"time"
)

type MobileEventsController struct {
	controller.BaseController
}

// URLMapping ...
func (c *MobileEventsController) URLMapping() {
	c.Mapping("ClockIn", c.ClockIn)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetRecent", c.GetRecent)
	// c.Mapping("Put", c.Put)
}

func (c *MobileEventsController) ClockIn() {
	if c.Verified {
		eventIdStr := c.Ctx.Input.Param(":id")
		eventId, _ := strconv.Atoi(eventIdStr)
		var currentDate = time.Now().Local()

		if found := models.FindEventClockByUserIDandEventID(c.Id, eventId); found {
			// already clocked in, return 200
			log.Println("Already sign up")
			c.Ctx.Output.SetStatus(200)
			if err := c.RenewToken(); err != nil {
				c.ServerError(err)
			}
		} else {
			// log.Println(currentDate)
			// need to add check if clock is already exist
			// for some reason the time add to database it will use UTC time to store???

			if _, err := models.AddEventClock(&models.EventClock{EventID: eventId, UserID: c.Id, StartTimeStamp: currentDate}); err == nil {
				c.Ctx.Output.SetStatus(200)
			} else {
				c.ServerError(err)
			}
			if err := c.RenewToken(); err != nil {
				c.ServerError(err)
			}
		}
	} else {
		c.Unauthorized()
	}
	c.ServeJSON()

}

func (c *MobileEventsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEventsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *MobileEventsController) GetRecent() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)

	// ----------------------------
	// modify
	// how many event should return
	var limit int64 = 0

	var currentDate = time.Now().Local()
	var beforeDate = currentDate.AddDate(0, 0, 10)
	var afterDate = currentDate.AddDate(0, 0, -2)

	var formatedBefore = fmt.Sprintf("%d-%02d-%02d", beforeDate.Year(), beforeDate.Month(), beforeDate.Day())
	var formatedAfter = fmt.Sprintf("%d-%02d-%02d", afterDate.Year(), afterDate.Month(), afterDate.Day())

	sortby = append(sortby, "startDate")
	order = append(order, "asc")
	// ----------------------------

	var offset int64
	// log.Println(c.Ctx.Request.Method)

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetRecentEvents(query, fields, sortby, order, offset, limit, formatedAfter, formatedBefore)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}
