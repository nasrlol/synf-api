package models

type UserInformation struct {
	UserID       uint8  `json:"id"`
	UserName     string `json:"name"`
	UserEmail    string `json:"email"`
	UserPassword string `json:"password"`
	UserRole     string `json:"role"`
	UserVerified bool   `json:"verified"`
}
