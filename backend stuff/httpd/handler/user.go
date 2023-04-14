package handler

import (
	"encoding/json"
	"fmt"
	"go-goal/util"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username    string // test this
	FirstName   string
	LastName    string
	Email       string
	Password    string
	XP          int
	Description string
}

// checks if the username is good. add more rules later
func IsValidUsername(globalDB *gorm.DB, username string) (exists bool, validName bool) {
	// Empty Username is not valid and does not exist
	if username == "" {
		return false, false
	}
	//Checks if the username is registered with another user
	globalDB.Model(&User{}).Select("count(*) > 0").Where("username = ?", username).Find(&exists)
	validName = !exists

	return exists, validName
}

// checks if the username is good. add more rules later
func IsValidEmail(globalDB *gorm.DB, email string) (exists bool, validEmail bool) {
	// Empty Email is not valid and does not exist
	if email == "" {
		return false, false
	}
	//Checks if the email is registered with another user
	globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exists)
	validEmail = !exists

	return exists, validEmail
}

// input json must contain all information of the user
func CreateUser(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var ThisUser User
		util.DecodeJSONRequest(&ThisUser, r.Body, w)

		returnInfo := struct { // this json return must be standardized
			Successful    bool
			ErrorExist    bool
			EmailExist    bool
			UsernameExist bool
		}{}

		// New User has 0 XP and a blank description
		ThisUser.XP = 0
		ThisUser.Description = ""

		returnInfo.EmailExist, _ = IsValidEmail(globalDB, ThisUser.Email)
		returnInfo.UsernameExist, _ = IsValidUsername(globalDB, ThisUser.Username)

		if !returnInfo.EmailExist && !returnInfo.UsernameExist {
			err := globalDB.Create(&ThisUser).Error
			if err != nil {
				returnInfo.ErrorExist = true
				json.NewEncoder(w).Encode(returnInfo)
				fmt.Println("Error in CreateUser")
				fmt.Printf("%+v\n", ThisUser)
				return
			}
			returnInfo.Successful = true
		} else {
			returnInfo.Successful = false
		}
		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input URL must have id of user
func GetUser(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		otherID := r.URL.Query().Get("id")
		returnInfo := struct {
			ThisUser   User
			ErrorExist bool
		}{}
		err := globalDB.Model(&User{}).First(&returnInfo.ThisUser, otherID).Error
		if err != nil {
			returnInfo.ErrorExist = true
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateUsername(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		returnInfo := struct { // Don't need to pass new or old information, should already have it
			ErrorExist    bool
			Successful    bool
			UsernameExist bool
			UsernameValid bool
		}{}

		params := mux.Vars(r)
		ID := params["id"]

		type NewUsername struct {
			Username string
		}
		var update NewUsername
		util.DecodeJSONRequest(&update, r.Body, w)

		returnInfo.UsernameExist, returnInfo.UsernameValid = IsValidUsername(globalDB, update.Username)

		if returnInfo.UsernameExist {
			returnInfo.ErrorExist = true
			returnInfo.Successful = false
			fmt.Printf("Error: Username already exists.")
			json.NewEncoder(w).Encode(returnInfo)
			return
		}
		if !returnInfo.UsernameValid {
			returnInfo.ErrorExist = true
			returnInfo.Successful = false
			fmt.Printf("Error: Username not valid.")
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		var user User
		err := globalDB.Model(&User{}).First(&user, ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%s", ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		globalDB.Model(&user).Update("username", update.Username)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateFirstname(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		type NewFirstname struct {
			Firstname string
		}
		var update NewFirstname
		util.DecodeJSONRequest(&update, r.Body, w)

		returnInfo := struct { // Don't need to pass new or old information, should already have it
			ErrorExist bool
			Successful bool
		}{}

		var user User
		err := globalDB.Model(&User{}).First(&user, ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%s", ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		globalDB.Model(&user).Update("first_name", update.Firstname)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateLastname(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		type NewLastname struct {
			Lastname string
		}
		var update NewLastname
		util.DecodeJSONRequest(&update, r.Body, w)
		returnInfo := struct { // Don't need to pass new or old information, should already have it
			ErrorExist bool
			Successful bool
		}{}

		var user User
		err := globalDB.Model(&User{}).First(&user, ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%s", ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		globalDB.Model(&user).Update("last_name", update.Lastname)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateEmail(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		returnInfo := struct { // Don't need to pass new or old information, should already have it
			ErrorExist bool
			Successful bool
			EmailExist bool
			EmailValid bool
		}{}

		params := mux.Vars(r)
		ID := params["id"]

		type NewEmail struct {
			Email string
		}
		var update NewEmail
		util.DecodeJSONRequest(&update, r.Body, w)

		returnInfo.EmailExist, returnInfo.EmailValid = IsValidEmail(globalDB, update.Email)

		if returnInfo.EmailExist {
			returnInfo.ErrorExist = true
			returnInfo.Successful = false
			fmt.Printf("Error: Email already exists.")
			json.NewEncoder(w).Encode(returnInfo)
			return
		}
		if !returnInfo.EmailValid {
			returnInfo.ErrorExist = true
			returnInfo.Successful = false
			fmt.Printf("Error: Email not valid.")
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		var user User
		err := globalDB.Model(&User{}).First(&user, ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%s", ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		globalDB.Model(&user).Update("email", update.Email)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdatePassword(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		type NewPassword struct {
			Password string
		}
		var update NewPassword
		util.DecodeJSONRequest(&update, r.Body, w)
		returnInfo := struct { // Don't need to pass new or old information, should already have it
			ErrorExist bool
			Successful bool
		}{}

		var user User
		err := globalDB.Model(&User{}).First(&user, ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%s", ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		globalDB.Model(&user).Update("password", update.Password)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func AddXP(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		type AdditionalXP struct {
			NewXP int
		}
		var update AdditionalXP
		util.DecodeJSONRequest(&update, r.Body, w)
		returnInfo := struct { // Don't need to pass new or old information, should already have it
			ErrorExist bool
			Successful bool
		}{}

		var user User
		err := globalDB.Model(&User{}).First(&user, ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%s", ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		update.NewXP += user.XP

		globalDB.Model(&user).Update("xp", update.NewXP)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateDescription(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		type NewDescription struct {
			Description string
		}
		var update NewDescription
		util.DecodeJSONRequest(&update, r.Body, w)
		returnInfo := struct { // Don't need to pass new or old information, should already have it
			ErrorExist bool
			Successful bool
		}{}

		var user User
		err := globalDB.Model(&User{}).First(&user, ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%s", ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

		globalDB.Model(&user).Update("description", update.Description)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

// must pass in json with attributes "Email" and "Password"
// returns a struct of whether email and password exists and a user object
// if both  email and password exists, the user object will be the corresponding user
// if not, the user object will have default values. the ID will be 0
func CheckLogin(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		returnInfo := struct { // the names of this json return must be standardized
			FindEmail    bool
			FindPassword bool
			ThisUser     User
		}{}

		params := mux.Vars(r)
		email := params["email"]
		password := params["password"]

		globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ?", email).Find(&returnInfo.FindEmail)
		globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ? AND password = ?", email, password).Find(&returnInfo.FindPassword)

		if returnInfo.FindEmail && returnInfo.FindPassword {
			globalDB.Where("email = ?", email).First(&returnInfo.ThisUser)
		}
		json.NewEncoder(w).Encode(returnInfo)
	}
}

// the URL must contain the username being checked
func CheckUsername(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		returnInfo := struct {
			Exists    bool
			ValidName bool
		}{}
		params := mux.Vars(r)
		username := params["username"]

		returnInfo.Exists, returnInfo.ValidName = IsValidUsername(globalDB, username)

		json.NewEncoder(w).Encode(returnInfo)
	}
}
