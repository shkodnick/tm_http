package profile

import (
	"app/internal/models"
)
type ProfileRepository interface {
	UserExists(email string) (bool, error)
	CreateUser(user *models.SignUpParams) error
	// SignUp
	// Login
	// Logout
}
