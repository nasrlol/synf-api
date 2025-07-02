package models

type DiskInformation struct {
	DiskID    int    `json:"id"`
	DiskName  string `json:"name"`
	DiskSpeed int    `json:"size"`
}
