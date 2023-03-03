package handler

import (
	"encoding/json"
	"fmt"
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

// input json must contain all information of the user
func CreateUser(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		returnInfo := struct { // this json return must be standardized
			Successful bool
			ErrorExist bool
			EmailExist bool
		}{}

		err := globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ?", user.Email).Find(&returnInfo.EmailExist).Error
		if err != nil {
			returnInfo.ErrorExist = true
		}

		if !returnInfo.EmailExist && !returnInfo.ErrorExist {
			r := globalDB.Create(&user)
			if r.Error != nil {
				returnInfo.ErrorExist = true
				json.NewEncoder(w).Encode(returnInfo)
				fmt.Println("Error in CreateUser")
				fmt.Println(user)
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
		json.NewDecoder(r.Body).Decode(&newUsername)
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
		json.NewDecoder(r.Body).Decode(&newFirstname)
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
		json.NewDecoder(r.Body).Decode(&newLastname)
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
		json.NewDecoder(r.Body).Decode(&newEmail)
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
		json.NewDecoder(r.Body).Decode(&newPassword)
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

// the structure of UpdateUser needs to be discussed, currently this is not working
// input json must contain id of user and all of the fields
func UpdateUser(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		returnInfo := struct {
			ErrorExist bool
			Successful bool
			ThisUser   User
		}{}

		var placeHolderUser User
		err := globalDB.Model(&user).Find(placeHolderUser, returnInfo.ThisUser.ID).Error
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in update user\nCould not find user with id:%d", returnInfo.ThisUser.ID)
			json.NewEncoder(w).Encode(returnInfo)
			return
		}

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

		json.NewDecoder(r.Body).Decode(&emailAndPassword)

		globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ?", emailAndPassword.Email).Find(&returnInfo.FindEmail)
		globalDB.Model(&User{}).Select("count(*) > 0").Where("email = ? AND password = ?", emailAndPassword.Email, emailAndPassword.Password).Find(&returnInfo.FindPassword)

		if returnInfo.FindEmail && returnInfo.FindPassword {
			globalDB.Where("email = ?", emailAndPassword.Email).First(&returnInfo.ThisUser)
		}
		json.NewEncoder(w).Encode(returnInfo)
	}
}
