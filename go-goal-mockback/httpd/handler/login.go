package handler

import (
	"encoding/json"
	"go-goal/platform/user"
	"net/http"
)

type loginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginOutput struct {
	LoggedIn  bool   `json:"loggedIn"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CheckLogin(G *user.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		accounts := G.GetAll()
		var loginInfo loginInput
		var response loginOutput

		json.NewDecoder(r.Body).Decode(&loginInfo)
		loginUser := LoginCheck(loginInfo, accounts)

		if loginUser == nil {
			response = loginOutput{
				false,
				"", "", "", "", "",
			}
		} else {
			response = loginOutput{
				true,
				loginUser.Username,
				loginUser.FirstName,
				loginUser.LastName,
				loginUser.Email,
				loginUser.Password,
			}
		}
		json.NewEncoder(w).Encode(response)
	}
}

// Returns a user if the login information is correct, nil otherwise
func LoginCheck(loginInfo loginInput, a []user.User) *user.User {
	for i := range a {
		if a[i].Username == loginInfo.Username &&
			a[i].Password == loginInfo.Password {
			// Repeat username or email
			return &a[i]
		}
	}
	return nil
}
