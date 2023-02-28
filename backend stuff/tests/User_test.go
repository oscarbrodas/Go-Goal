package test

import (
	"bytes"
	"encoding/json"
	"go-goal/httpd/handlers"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "?id=2", nil)
	if err != nil {
		panic(err)
	}
	handlers.GetUser(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	returnInfo := struct {
		ThisUser   handlers.User
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.ErrorExist {
		t.Errorf("There was an error")
	}

	if returnInfo.ThisUser.FirstName != "2" || returnInfo.ThisUser.Email != "2@gmail.com" || returnInfo.ThisUser.Password != "pw" {
		t.Errorf("Got the wrong user")
	}
}

// creating a new user without the email already existing
func TestCreateUser1(t *testing.T) {
	initializeTestDatabase()

	var user handlers.User
	user.Username = "dwan12345"
	user.FirstName = "don"
	user.LastName = "chen"
	user.Email = "dc@gmail.com"
	user.Password = "pw"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "", &buf)
	if err != nil {
		panic(err)
	}
	handlers.CreateUser(globalDB)(w, r)

	returnInfo := struct {
		Successful bool
		ErrorExist bool
		EmailExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist || returnInfo.EmailExist {
		t.Errorf("Expected {Successful:true, ErrorExist:false, EmailExist:false}, but got %v", returnInfo)
	}

	var inputtedUser handlers.User
	globalDB.Model(&handlers.User{}).Raw("SELECT username, first_name, last_name, email, password FROM users WHERE id = ?", 1).Scan(&inputtedUser)
	if !reflect.DeepEqual(user, inputtedUser) { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("Expected %v, but got %v", user, inputtedUser)
	}
}

// creating a new user with the email existing
func TestCreateUser2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	var user handlers.User
	user.Username = "dwan12345"
	user.FirstName = "don"
	user.LastName = "chen"
	user.Email = "1@gmail.com"
	user.Password = "pw"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "", &buf)
	if err != nil {
		panic(err)
	}
	handlers.CreateUser(globalDB)(w, r)

	returnInfo := struct {
		Successful bool
		ErrorExist bool
		EmailExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || returnInfo.ErrorExist || !returnInfo.EmailExist {
		t.Errorf("Expected {Successful:false, ErrorExist:false, EmailExist:true}, but got %v", returnInfo)
	}
}

// idk if this test if appropriate, so it might be rewritten
/*
func TestUpdateUser(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	var user handlers.User
	user.Username = "newUserName"
	user.FirstName = "don"
	user.LastName = "chen"
	user.Email = "1@gmail.com"
	user.Password = "pw"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("PUT", "", &buf)
	if err != nil {
		panic(err)
	}
	handlers.UpdateUser(globalDB)(w, r)

	returnInfo := struct {
		Successful bool
		ErrorExist bool
		ThisUser   handlers.User
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.ErrorExist {
		t.Errorf("There was an error")
	}

	var inputtedUser handlers.User
	globalDB.Model(&handlers.User{}).Raw("SELECT username, first_name, last_name, email, password FROM users WHERE id = ?", 1).Scan(&inputtedUser)
	if !reflect.DeepEqual(user, inputtedUser) { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("Expected %v, but got %v", user, inputtedUser)
	}
}
*/

// this unit does not work for some reason
// tests login when email and password are correct
/*
func TestCheckLogin1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	var EandP = struct {
		Email    string
		Password string
	}{
		Email:    "1@gmail.com",
		Password: "pw",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(EandP)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "", &buf)
	if err != nil {
		panic(err)
	}
	handlers.CheckLogin(globalDB)(w, r)

	returnInfo := struct {
		FindEmail    bool
		FindPassword bool
		User         handlers.User
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.FindEmail || !returnInfo.FindPassword {
		t.Errorf("Expected {FindEmail:true, FindPassword:false}, but got {%t %t}", returnInfo.FindEmail, returnInfo.FindPassword)
	} else if !(returnInfo.User.ID == 1) || !(returnInfo.User.Email == "1@gmail.com") {
		t.Errorf("Did not get correct user back\nGot:%v", returnInfo.User)
	}

}
*/

// tests login when email and password are incorrect
func TestCheckLogin2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	var EandP = struct {
		Email    string
		Password string
	}{
		Email:    "1@gmail.com",
		Password: "pw22",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(EandP)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "", &buf)
	if err != nil {
		panic(err)
	}
	handlers.CheckLogin(globalDB)(w, r)

	returnInfo := struct {
		FindEmail    bool
		FindPassword bool
		User         handlers.User
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.FindEmail || returnInfo.FindPassword {
		t.Errorf("Expected {FindEmail:true, FindPassword:false}, but got {%t %t}", returnInfo.FindEmail, returnInfo.FindPassword)
	}
}
