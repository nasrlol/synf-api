package gpu

import (

)

type GpuInformation struct{
	CpuID	uint8	`json:"cpu_id"`
	CpuName	string	`json:"cpu_name"`
	CpuTemp	uint8	`json:"cpu_temp"`
	CpuFreq	uint8	`json:"cpu_freq"`
}



