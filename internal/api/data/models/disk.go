package models

type Disk struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Size int    `json:"size"`
}
