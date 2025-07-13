package models

type Cpu struct {
	Id   uint8  `json:"cpu_id"`
	Name string `json:"cpu_name"`
	Temp uint8  `json:"cpu_temp"`
	Freq uint8  `json:"cpu_freq"`
}
