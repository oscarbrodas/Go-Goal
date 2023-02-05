package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string
	Password  string
}

// checks if email exists, if not then creates the user
// returns json of whether email was found
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	emailExist := struct { // this json return must be standardized
		FindEmail bool
	}{}

	c := int64(0)
	globalDB.Model(&User{}).Where("email = ?", user.Email).Count(&c)
	if c > 0 {
		emailExist.FindEmail = true
	}

	json.NewEncoder(w).Encode(emailExist)
	if !emailExist.FindEmail {
		globalDB.Create(&user)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user []User
	globalDB.Find(&user)
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	params := mux.Vars(r)
	globalDB.First(&user, params["id"]) //grabs {id} from r.handleFunc in main

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(r)
	globalDB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	globalDB.Save(&user)

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(r)
	globalDB.Delete(&user, params["id"])

	json.NewEncoder(w).Encode(user)
}

// must pass in json with attributes "Email" and "Password"
// returns a struct of whether email and password exists
// if both exist, return the user object
func checkLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	// these create temporary structs
	emailAndPassword := struct {
		Email    string
		Password string
	}{}
	returnInfo := struct { // the names of this json return must be standardized
		FindEmail    bool
		FindPassword bool
	}{}

	json.NewDecoder(r.Body).Decode(&emailAndPassword)

	c := int64(0)
	globalDB.Model(&User{}).Where("email = ?", emailAndPassword.Email).Count(&c)
	if c > 0 {
		returnInfo.FindEmail = true
	}
	globalDB.Model(&User{}).Where("email = ? AND password = ?", emailAndPassword.Email, emailAndPassword.Password).Count(&c)
	if c > 0 {
		returnInfo.FindPassword = true
	}

	// always returns returnInfo, only returns user if email and password are correct
	json.NewEncoder(w).Encode(returnInfo)
	if returnInfo.FindEmail && returnInfo.FindPassword {
		globalDB.Where("email = ?", emailAndPassword.Email).Find(&user)
		json.NewEncoder(w).Encode(user)
	}
}
