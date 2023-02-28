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

// this function is not neccessary. the route is removed
func GetUsers(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user []User
		globalDB.Find(&user)
		json.NewEncoder(w).Encode(user)
	}
}

// input URL must has id of user
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

// delete user is complicated. the route is removed
func DeleteUser(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user User
		params := mux.Vars(r)
		globalDB.Delete(&user, params["id"])

		json.NewEncoder(w).Encode(user)
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
