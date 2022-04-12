package models

import (
	"net/http"
)

type DeviceInfo struct {
	Name  string `json:"Name"`
	Model string `json:"Model"`
	UUID  string `json:"UUID"`
}

func (m *DeviceInfo) ConstructFromHeader(header http.Header) {
	m.Name = header.Get("DeviceName")
	m.UUID = header.Get("DeviceUUID")
	m.Model = header.Get("DeviceModel")
}
