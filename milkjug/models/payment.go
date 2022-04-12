package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Payment struct {
	AccountType     int       `orm:"column(accountType);null"`
	Location        string    `orm:"column(location);size(30);null"`
	Id              int       `orm:"column(paymentID);auto"`
	PaymentMethod   int       `orm:"column(paymentMethod);null"`
	PaymentPlanType int       `orm:"column(paymentPlanType);null"`
	ProcessDate     time.Time `orm:"column(processDate);type(datetime);null"`
	Semester        int       `orm:"column(semester);null"`
	TotalBalance    float32   `orm:"column(totalBalance);null"`
	UserID          int       `orm:"column(userID);null"`
}

func (t *Payment) TableName() string {
	return "payment"
}

func init() {
	orm.RegisterModel(new(Payment))
}

// AddPayment insert a new Payment into database and returns
// last inserted Id on success.
func AddPayment(m *Payment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPaymentById retrieves Payment by Id. Returns error if
// Id doesn't exist
func GetPaymentById(id int) (v *Payment, err error) {
	o := orm.NewOrm()
	v = &Payment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllAccountPayment(userID int) ([]Payment, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Payment))
	qs = qs.Filter("userID", userID)

	var sortFields []string
	sortFields = append(sortFields, "AccountType")
	sortFields = append(sortFields, "-ProcessDate")
	qs = qs.OrderBy(sortFields...)

	var l []Payment
	_, err := qs.All(&l)
	return l, err
}

// GetAllPayment retrieves all Payment matches certain condition. Returns empty list if
// no records exist
func GetAllPayment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Payment))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Payment
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePayment updates Payment by Id and returns error if
// the record to be updated doesn't exist
func UpdatePaymentById(m *Payment) (err error) {
	o := orm.NewOrm()
	v := Payment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePayment deletes Payment by Id and returns error if
// the record to be deleted doesn't exist
func DeletePayment(id int) (err error) {
	o := orm.NewOrm()
	v := Payment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Payment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
