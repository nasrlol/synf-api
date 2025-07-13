package models

type Disk struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Speed int    `json:"size"`
}
