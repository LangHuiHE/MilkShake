package models

import (
	"encoding/json"
	"strconv"
)

type LoginWithStringID struct {
	Id           string `json:"Id"`
	UserPassword string `json:"UserPassword"`
}

type LoginTemp struct {
	Id           int    `json:"Id"`
	UserPassword string `json:"UserPassword"`
}

func (m *LoginTemp) SensitizeLogin(rawData []byte) error {
	var s LoginWithStringID
	if json.Unmarshal(rawData, &s) == nil {
		if id, err := strconv.Atoi(s.Id); err == nil {
			m.Id = id
			m.UserPassword = s.UserPassword
			return nil
		} else {
			return err
		}
	} else {
		return json.Unmarshal(rawData, m)
	}
}
