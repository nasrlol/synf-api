package models

type RamInformation struct {
	RamID   uint8  `json:"ram_id"`
	RamName string `json:"ram_name"`
	RamTemp uint8  `json:"ram_temp"`
	RamFreq uint8  `json:"ram_freq"`
}
