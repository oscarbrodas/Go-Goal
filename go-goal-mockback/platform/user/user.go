package user

import "gorm.io/gorm"

// User data
type User struct {
	gorm.Model
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
