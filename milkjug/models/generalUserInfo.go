package models

import (
	"time"
)

type GeneralUserInfo struct {
	ContactAddress   string
	EmailAddress     string
	EnrollDate       time.Time
	FirstName        string
	LastName         string
	PermanentAddress string
	PhoneNumber      string
	Id               int
	UserType         int
	StudentInfo      StudentInfo
}

type StudentInfo struct {
	Credits    int
	Department string
	Major      string
}

func (c *GeneralUserInfo) CreateGeneralUserInfo(id int) error {
	if u, err := GetUsersById(id); err == nil {
		c.ContactAddress = u.ContactAddress
		c.EmailAddress = u.EmailAddress
		c.EnrollDate = u.EnrollDate
		c.FirstName = u.FirstName
		c.LastName = u.LastName
		c.PermanentAddress = u.PermanentAddress
		c.PhoneNumber = u.PhoneNumber
		c.Id = u.Id
		c.UserType = u.UserType
		if c.UserType == 3 {
			if err = c.AddStudentInfo(); err != nil {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}

func (c *GeneralUserInfo) AddStudentInfo() error {
	info, err := GetStudentsById(c.Id)
	if err != nil {
		return err
	}
	c.StudentInfo.Credits = info.Credits

	d, err := GetStudentDepartmentById(info.Department)
	if err != nil {
		return err
	}
	c.StudentInfo.Department = d.Definition
	
	m, err := GetStudentMajorById(info.Major)
	c.StudentInfo.Major = m.Definition
	return err
}
