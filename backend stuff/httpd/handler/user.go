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

	Username  string // test this
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// checks if the username is good. add more rules later
func isValidUsername(globalDB *gorm.DB, username string) (exists bool, validName bool) {
	globalDB.Model(&User{}).Select("count(*) > 0").Where("username = ?", username).Find(&exists)
	validName = !exists

	if username == "" {
		validName = false
	}

	return exists, validName
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

		err := globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ?", ThisUser.Email).Find(&returnInfo.EmailExist).Error
		if err != nil {
			returnInfo.ErrorExist = true
		}
		exists, _ := isValidUsername(globalDB, ThisUser.Username)
		returnInfo.UsernameExist = exists

		if !returnInfo.EmailExist && !returnInfo.ErrorExist && !returnInfo.UsernameExist {
			err = globalDB.Create(&ThisUser).Error
			if err != nil {
				returnInfo.ErrorExist = true
				json.NewEncoder(w).Encode(returnInfo)
				fmt.Println("Error in CreateUser")
				fmt.Printf("%+v\n", ThisUser)
				return
			}
			returnInfo.Successful = true
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
		params := mux.Vars(r)
		ID := params["id"]

		var newUsername string
		util.DecodeJSONRequest(&newUsername, r.Body, w)
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

		globalDB.Model(&user).Update("username", newUsername)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateFirstname(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		var newFirstname string
		util.DecodeJSONRequest(&newFirstname, r.Body, w)

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

		globalDB.Model(&user).Update("first_name", newFirstname)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateLastname(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		var newLastname string
		util.DecodeJSONRequest(&newLastname, r.Body, w)
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

		globalDB.Model(&user).Update("last_name", newLastname)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateEmail(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		var newEmail string
		util.DecodeJSONRequest(&newEmail, r.Body, w)
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

		globalDB.Model(&user).Update("email", newEmail)

		returnInfo.Successful = true
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdatePassword(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]

		var newPassword string
		util.DecodeJSONRequest(&newPassword, r.Body, w)
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

		globalDB.Model(&user).Update("password", newPassword)

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

		// these create temporary structs
		emailAndPassword := struct {
			Email    string
			Password string
		}{}
		returnInfo := struct { // the names of this json return must be standardized
			FindEmail    bool
			FindPassword bool
			ThisUser     User
		}{}

		util.DecodeJSONRequest(&emailAndPassword, r.Body, w)

		globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ?", emailAndPassword.Email).Find(&returnInfo.FindEmail)
		globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ? AND password = ?", emailAndPassword.Email, emailAndPassword.Password).Find(&returnInfo.FindPassword)

		if returnInfo.FindEmail && returnInfo.FindPassword {
			globalDB.Where("email = ?", emailAndPassword.Email).First(&returnInfo.ThisUser)
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

		returnInfo.Exists, returnInfo.ValidName = isValidUsername(globalDB, username)

		json.NewEncoder(w).Encode(returnInfo)
	}
}
