package handler

import (
	"encoding/json"
	"go-goal/platform/user"
	"net/http"

	"gorm.io/gorm"
)

func SignUp(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		newUser := user.User{}
		json.NewDecoder(r.Body).Decode(&newUser)
		var response bool
		var accounts []user.User
		db.Find(&accounts)
		if FoundUser(newUser, accounts) {
			response = false
		} else {
			response = true
			db.Create(&newUser)
		}
		json.NewEncoder(w).Encode(response)
	}
}

// Returns true if the user already exsits
func FoundUser(user user.User, a []user.User) bool {
	for i := range a {
		if a[i].Username == user.Username ||
			a[i].Email == user.Email {
			// Repeat username or email
			return true
		}
	}
	return false
}
