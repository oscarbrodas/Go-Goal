package handler

import (
	"encoding/json"
	"go-goal/platform/user"
	"net/http"
)

func SignUp(A *user.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		newUser := user.User{}
		json.NewDecoder(r.Body).Decode(&newUser)
		var response bool
		accounts := A.GetAll()
		if FoundUser(newUser, accounts) {
			response = false
		} else {
			response = true
			A.Add(newUser)
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
