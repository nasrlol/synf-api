package models

type User struct {
	Id       uint8  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Verified bool   `json:"verified"`
}
