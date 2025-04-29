package disk

import (

)

type DiskInformation struct{
	CpuID	uint8	`json:"disk_id"`
	CpuName	string	`json:"disk_name"`
	CpuTemp	uint8	`json:"disk_temp"`
	CpuFreq	uint8	`json:"disk_speed"`
} 


