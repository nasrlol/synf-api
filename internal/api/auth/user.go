package auth

import (
	"synf/internal/api/data/models"
)

func GetPolicy(user models.User) bool {
	if user.Role != "user" || user.Role != "admin" {
		return false
	}
	return true
}

func CreatePolicy(user models.User) bool {
	return true
}
