package models

type DeviceInformation struct {
	DeviceID uint8  `json:"device_id"`
	UpTime   string `json:"device_upTime"`
}
