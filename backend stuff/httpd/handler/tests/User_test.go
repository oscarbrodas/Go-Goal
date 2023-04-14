package handler_test

import (
	"bytes"
	"encoding/json"
	"go-goal/httpd/handler"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestIsValidUsername(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"test\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")

	success := "unique"
	fail := "test"

	existsTrue, isValidTrue := handler.IsValidUsername(globalDB, success)
	existsFalse, isValidFalse := handler.IsValidUsername(globalDB, fail)

	if existsTrue && isValidTrue {
		t.Errorf("Expected first passed username to be accepted: Exists: \"%t\" Valid: \"%t\"", existsTrue, isValidTrue)
	}

	if !existsTrue && !isValidTrue {
		t.Errorf("Expected second passed username to be denied: Exists: \"%t\" Valid: \"%t\"", existsFalse, isValidFalse)
	}
}

func TestIsValidEmail(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"test@gmail.com\",\"pw\")")

	success := "unique@gmail.com"
	fail := "test@gmail.com"

	existsTrue, isValidTrue := handler.IsValidEmail(globalDB, success)
	existsFalse, isValidFalse := handler.IsValidEmail(globalDB, fail)

	if existsTrue && isValidTrue {
		t.Errorf("Expected first passed email to be accepted: Exists: \"%t\" Valid: \"%t\"", existsTrue, isValidTrue)
	}

	if !existsTrue && !isValidTrue {
		t.Errorf("Expected second passed email to be denied: Exists: \"%t\" Valid: \"%t\"", existsFalse, isValidFalse)
	}
}

func TestGetUser(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username, first_name,last_name,email,password, xp, description) values(\"dc\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw1\",\"0\",\"Hi!\")")
	globalDB.Exec("insert into users(username, first_name,last_name,email,password, xp, description) values(\"or\",\"Oscar\",\"Rodas\",\"2@gmail.com\",\"pw2\",\"100\",\"Hi?\")")

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "?id=2", nil)
	if err != nil {
		panic(err)
	}
	handler.GetUser(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	returnInfo := struct {
		ThisUser   handler.User
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.ErrorExist {
		t.Errorf("There was an error")
	}

	if returnInfo.ThisUser.Username != "or" || returnInfo.ThisUser.FirstName != "Oscar" || returnInfo.ThisUser.LastName != "Rodas" || returnInfo.ThisUser.Email != "2@gmail.com" ||
		returnInfo.ThisUser.Password != "pw2" || returnInfo.ThisUser.XP != 100 || returnInfo.ThisUser.Description != "Hi?" {
		t.Errorf("Got the wrong user")
	}
}

// creating a new user without the email already existing
func TestCreateUser1(t *testing.T) {
	initializeTestDatabase()

	var user handler.User
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
	handler.CreateUser(globalDB)(w, r)

	returnInfo := struct {
		Successful    bool
		ErrorExist    bool
		EmailExist    bool
		UsernameExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist || returnInfo.EmailExist || returnInfo.UsernameExist {
		t.Errorf("Expected {Successful:true, ErrorExist:false, EmailExist:false, UsernameExist:false}, but got %v", returnInfo)
	}

	var inputtedUser handler.User
	globalDB.Model(&handler.User{}).Raw("SELECT username, first_name, last_name, email, password, xp, description FROM users WHERE id = ?", 1).Scan(&inputtedUser)
	if !reflect.DeepEqual(user, inputtedUser) { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("Expected %v, but got %v", user, inputtedUser)
	}
}

// creating a new user with the email existing
func TestCreateUser2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp, description) values(\"dc\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw1\",\"0\",\"Hi!\")")

	var user handler.User
	user.Username = "or"
	user.FirstName = "Oscar"
	user.LastName = "Rodas"
	user.Email = "1@gmail.com"
	user.Password = "pw2"
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
	handler.CreateUser(globalDB)(w, r)

	returnInfo := struct {
		Successful    bool
		ErrorExist    bool
		EmailExist    bool
		UsernameExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || returnInfo.ErrorExist || !returnInfo.EmailExist || returnInfo.UsernameExist {
		t.Errorf("Expected {Successful:false, ErrorExist:false, EmailExist:true, UsernameExist: false}, but got %v", returnInfo)
	}
}

// creating a new user with the username existing
func TestCreateUser3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp, description) values(\"dc\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw1\",\"0\",\"Hi!\")")

	var user handler.User
	user.Username = "dc"
	user.FirstName = "Oscar"
	user.LastName = "Rodas"
	user.Email = "2@gmail.com"
	user.Password = "pw2"
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
	handler.CreateUser(globalDB)(w, r)

	returnInfo := struct {
		Successful    bool
		ErrorExist    bool
		EmailExist    bool
		UsernameExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || returnInfo.ErrorExist || returnInfo.EmailExist || !returnInfo.UsernameExist {
		t.Errorf("Expected {Successful:false, ErrorExist:false, EmailExist:false, UsernameExist: true}, but got %v", returnInfo)
	}
}

// Updating a table with a single user, correct id given
func TestUpdateUsername1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"old\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")

	type NewUsername struct {
		Username string
	}
	var update NewUsername
	update.Username = "new"

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/1/username", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/username", handler.UpdateUsername(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist    bool
		Successful    bool
		UsernameExist bool
		UsernameValid bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.Username != update.Username {
		t.Errorf("Expected to update username to \"%s\", but it is \"%s\"", update.Username, user.Username)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateUsername2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"preserved\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"old\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\")")

	type NewUsername struct {
		Username string
	}
	var update NewUsername
	update.Username = "new"

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/2/username", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/username", handler.UpdateUsername(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var preservedUser handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser).First(&preservedUser, 1)
	globalDB.Model(&updatedUser).First(&updatedUser, 2)

	if preservedUser.Username != "preserved" {
		t.Errorf("Expected to preserve username as \"%s\", but it is \"%s\"", "preserved", preservedUser.Username)
	}

	if updatedUser.Username != update.Username {
		t.Errorf("Expected to update username to \"%s\", but it is \"%s\"", update.Username, updatedUser.Username)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateUsername3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"preserved\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"preserved\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\")")

	type NewUsername struct {
		Username string
	}
	var update NewUsername
	update.Username = "new"

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/3/username", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/username", handler.UpdateUsername(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || !returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True} , but got %v", returnInfo)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.Username != "preserved" {
		t.Errorf("Expected to preserve username as \"%s\", but it is \"%s\"", "preserved", preservedUser1.Username)
	}

	if preservedUser2.Username != "preserved" {
		t.Errorf("Expected to preserve username as \"%s\", but it is \"%s\"", "preserved", preservedUser2.Username)
	}

	err = globalDB.Model(&handler.User{}).First(&updatedUser, 3).Error
	if err == nil { // There should not be a third user
		t.Errorf("Expected error to exist, user exists: %v", updatedUser)
	}
}

// Updating a table with a two users, updating both users with the other's username, should not change the database
func TestUpdateUsername4(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC1\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC2\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\")")

	type NewUsername struct {
		Username string
	}
	var update1 NewUsername
	var update2 NewUsername

	update1.Username = "DC1"
	update2.Username = "DC2"

	var buf1 bytes.Buffer
	var buf2 bytes.Buffer

	err1 := json.NewEncoder(&buf1).Encode(update1)
	err2 := json.NewEncoder(&buf2).Encode(update2)

	req1, err1 := http.NewRequest("PUT", "/api/users/2/username", &buf1)
	if err1 != nil {
		t.Fatal(err1)
	}
	req2, err2 := http.NewRequest("PUT", "/api/users/1/username", &buf2)
	if err2 != nil {
		t.Fatal(err2)
	}

	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/username", handler.UpdateUsername(globalDB)).Methods("PUT")
	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)

	returnInfo1 := struct {
		ErrorExist    bool
		Successful    bool
		UsernameExist bool
		UsernameValid bool
	}{}
	returnInfo2 := struct {
		ErrorExist    bool
		Successful    bool
		UsernameExist bool
		UsernameValid bool
	}{}

	json.NewDecoder(w1.Result().Body).Decode(&returnInfo1)
	json.NewDecoder(w2.Result().Body).Decode(&returnInfo2)

	if returnInfo1.Successful || !returnInfo1.ErrorExist || !returnInfo1.UsernameExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True, UsernameExist:True} , but got %v", returnInfo1)
	}

	if returnInfo2.Successful || !returnInfo2.ErrorExist || !returnInfo2.UsernameExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True, UsernameExist:True} , but got %v", returnInfo1)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.Username != "DC1" {
		t.Errorf("Expected to preserve username as \"%s\", but it is \"%s\"", "DC1", preservedUser1.Username)
	}

	if preservedUser2.Username != "DC2" {
		t.Errorf("Expected to preserve username as \"%s\", but it is \"%s\"", "DC2", preservedUser2.Username)
	}
}

// Updating a table with a single user, correct id given
func TestUpdateFirstname1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"old\",\"Chen\",\"1@gmail.com\",\"pw\")")

	type NewFirstname struct {
		Firstname string
	}
	var update NewFirstname
	update.Firstname = "New"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/1/firstname", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/firstname", handler.UpdateFirstname(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.FirstName != update.Firstname {
		t.Errorf("Expected to update firstname to \"%s\", but it is \"%s\"", update.Firstname, user.FirstName)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateFirstname2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"preserved\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"old\", \"Chen\",\"2@gmail.com\",\"pw\")")

	type NewFirstname struct {
		Firstname string
	}
	var update NewFirstname
	update.Firstname = "New"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/2/firstname", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/firstname", handler.UpdateFirstname(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var preservedUser handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser).First(&preservedUser, 1)
	globalDB.Model(&updatedUser).First(&updatedUser, 2)

	if preservedUser.FirstName != "preserved" {
		t.Errorf("Expected to preserve firstname as \"%s\", but it is \"%s\"", "preserved", preservedUser.FirstName)
	}

	if updatedUser.FirstName != update.Firstname {
		t.Errorf("Expected to update firstname to \"%s\", but it is \"%s\"", update.Firstname, updatedUser.FirstName)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateFirstname3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"preserved\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"preserved\", \"Chen\",\"2@gmail.com\",\"pw\")")

	type NewFirstname struct {
		Firstname string
	}
	var update NewFirstname
	update.Firstname = "New"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/3/firstname", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/firstname", handler.UpdateFirstname(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || !returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True} , but got %v", returnInfo)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.FirstName != "preserved" {
		t.Errorf("Expected to preserve firstname as \"%s\", but it is \"%s\"", "preserved", preservedUser1.FirstName)
	}

	if preservedUser2.FirstName != "preserved" {
		t.Errorf("Expected to preserve firstname as \"%s\", but it is \"%s\"", "preserved", preservedUser2.FirstName)
	}

	err = globalDB.Model(&handler.User{}).First(&updatedUser, 3).Error
	if err == nil { // There should not be a third user
		t.Errorf("Expected error, but user exsists: %v", updatedUser)
	}
}

// Updating a table with a single user, correct id given
func TestUpdateLastname1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"old\",\"1@gmail.com\",\"pw\")")

	type NewLastname struct {
		Lastname string
	}
	var update NewLastname
	update.Lastname = "New"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/1/lastname", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/lastname", handler.UpdateLastname(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.LastName != update.Lastname {
		t.Errorf("Expected to update lastname to \"%s\", but it is \"%s\"", update.Lastname, user.LastName)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateLastname2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"preserved\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"old\",\"2@gmail.com\",\"pw\")")

	type NewLastname struct {
		Lastname string
	}
	var update NewLastname
	update.Lastname = "New"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/2/lastname", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/lastname", handler.UpdateLastname(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var preservedUser handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser).First(&preservedUser, 1)
	globalDB.Model(&updatedUser).First(&updatedUser, 2)

	if preservedUser.LastName != "preserved" {
		t.Errorf("Expected to preserve lastname as \"%s\", but it is \"%s\"", "preserved", preservedUser.LastName)
	}

	if updatedUser.LastName != update.Lastname {
		t.Errorf("Expected to update lastname to \"%s\", but it is \"%s\"", update.Lastname, updatedUser.LastName)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateLastname3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"preserved\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"preserved\",\"2@gmail.com\",\"pw\")")

	type NewLastname struct {
		Lastname string
	}
	var update NewLastname
	update.Lastname = "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/3/lastname", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/lastname", handler.UpdateLastname(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || !returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True} , but got %v", returnInfo)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.LastName != "preserved" {
		t.Errorf("Expected to preserve lastname as \"%s\", but it is \"%s\"", "preserved", preservedUser1.LastName)
	}

	if preservedUser2.LastName != "preserved" {
		t.Errorf("Expected to preserve lastname as \"%s\", but it is \"%s\"", "preserved", preservedUser2.LastName)
	}

	err = globalDB.Model(&handler.User{}).First(&updatedUser, 3).Error
	if err == nil { // There should not be a third user
		t.Errorf("Expected error, but user exsists: %v", updatedUser)
	}
}

// Updating a table with a single user, correct id given
func TestUpdateEmail1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"old@gmail.com\",\"pw\")")

	type NewEmail struct {
		Email string
	}
	var update NewEmail
	update.Email = "new@gmail.com"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/1/email", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/email", handler.UpdateEmail(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.Email != update.Email {
		t.Errorf("Expected to update email to \"%s\", but it is \"%s\"", update.Email, user.Email)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateEmail2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"preserved@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"old@gmail.com\",\"pw\")")

	type NewEmail struct {
		Email string
	}
	var update NewEmail
	update.Email = "new@gmail.com"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/2/email", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/email", handler.UpdateEmail(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var preservedUser handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser).First(&preservedUser, 1)
	globalDB.Model(&updatedUser).First(&updatedUser, 2)

	if preservedUser.Email != "preserved@gmail.com" {
		t.Errorf("Expected to preserve email as \"%s\", but it is \"%s\"", "preserved", preservedUser.Email)
	}

	if updatedUser.Email != update.Email {
		t.Errorf("Expected to update eamil to \"%s\", but it is \"%s\"", update.Email, updatedUser.Email)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateEmail3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"preserved1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"preserved2@gmail.com\",\"pw\")")

	type NewEmail struct {
		Email string
	}
	var update NewEmail
	update.Email = "new@gmail.com"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/3/email", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/email", handler.UpdateEmail(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || !returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True} , but got %v", returnInfo)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.Email != "preserved1@gmail.com" {
		t.Errorf("Expected to preserve email as \"%s\", but it is \"%s\"", "preserved", preservedUser1.Email)
	}

	if preservedUser2.Email != "preserved2@gmail.com" {
		t.Errorf("Expected to preserve email as \"%s\", but it is \"%s\"", "preserved", preservedUser2.Email)
	}

	err = globalDB.Model(&handler.User{}).First(&updatedUser, 3).Error
	if err == nil { // There should not be a third user
		t.Errorf("Expected error, but user exsists: %v", updatedUser)
	}
}

func TestUpdateEmail4(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC1\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC2\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\")")

	type NewEmail struct {
		Email string
	}
	var update1 NewEmail
	var update2 NewEmail

	update1.Email = "1@gmail.com"
	update2.Email = "2@gmail.com"

	var buf1 bytes.Buffer
	var buf2 bytes.Buffer

	err1 := json.NewEncoder(&buf1).Encode(update1)
	err2 := json.NewEncoder(&buf2).Encode(update2)

	req1, err1 := http.NewRequest("PUT", "/api/users/2/email", &buf1)
	if err1 != nil {
		t.Fatal(err1)
	}
	req2, err2 := http.NewRequest("PUT", "/api/users/1/email", &buf2)
	if err2 != nil {
		t.Fatal(err2)
	}

	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/email", handler.UpdateEmail(globalDB)).Methods("PUT")
	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)

	returnInfo1 := struct {
		ErrorExist bool
		Successful bool
		EmailExist bool
		EmailValid bool
	}{}
	returnInfo2 := struct {
		ErrorExist bool
		Successful bool
		EmailExist bool
		EmailValid bool
	}{}

	json.NewDecoder(w1.Result().Body).Decode(&returnInfo1)
	json.NewDecoder(w2.Result().Body).Decode(&returnInfo2)

	if returnInfo1.Successful || !returnInfo1.ErrorExist || !returnInfo1.EmailExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True, UsernameExist:True} , but got %v", returnInfo1)
	}

	if returnInfo2.Successful || !returnInfo2.ErrorExist || !returnInfo2.EmailExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True, UsernameExist:True} , but got %v", returnInfo1)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.Email != "1@gmail.com" {
		t.Errorf("Expected to preserve email as \"%s\", but it is \"%s\"", "1@gmail.com", preservedUser1.Email)
	}

	if preservedUser2.Email != "2@gmail.com" {
		t.Errorf("Expected to preserve email as \"%s\", but it is \"%s\"", "2@gmail.com", preservedUser2.Username)
	}
}

// Updating a table with a single user, correct id given
func TestUpdatePassword1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"old\")")

	type NewPassword struct {
		Password string
	}
	var update NewPassword
	update.Password = "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/1/password", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/password", handler.UpdatePassword(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.Password != update.Password {
		t.Errorf("Expected to update password to \"%s\", but it is \"%s\"", update.Password, user.Password)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdatePassword2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"preserved\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"2@gmail.com\",\"old\")")

	type NewPassword struct {
		Password string
	}
	var update NewPassword
	update.Password = "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/2/password", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/password", handler.UpdatePassword(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var preservedUser handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser).First(&preservedUser, 1)
	globalDB.Model(&updatedUser).First(&updatedUser, 2)

	if preservedUser.Password != "preserved" {
		t.Errorf("Expected to preserve password as \"%s\", but it is \"%s\"", "preserved", preservedUser.Password)
	}

	if updatedUser.Password != update.Password {
		t.Errorf("Expected to update password to \"%s\", but it is \"%s\"", update.Password, updatedUser.Password)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdatePassword3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"preserved\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"2@gmail.com\",\"preserved\")")

	type NewPassword struct {
		Password string
	}
	var update NewPassword
	update.Password = "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/3/password", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/password", handler.UpdatePassword(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || !returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True} , but got %v", returnInfo)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.Password != "preserved" {
		t.Errorf("Expected to preserve password as \"%s\", but it is \"%s\"", "preserved", preservedUser1.Password)
	}

	if preservedUser2.Password != "preserved" {
		t.Errorf("Expected to preserve password as \"%s\", but it is \"%s\"", "preserved", preservedUser2.Password)
	}

	err = globalDB.Model(&handler.User{}).First(&updatedUser, 3).Error
	if err == nil { // There should not be a third user
		t.Errorf("Expected error, but user exsists: %v", updatedUser)
	}
}

func TestAddXP(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"old\", \"50\")")

	type AdditionalXP struct {
		NewXP int
	}
	var update AdditionalXP
	update.NewXP = 50
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/1/xp", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/xp", handler.AddXP(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.XP != 100 {
		t.Errorf("Expected to update xp to \"%d\", but it is \"%d\"", 100, user.XP)
	}
}

// Updating a table with a single user, correct id given
func TestUpdateDescription1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp,description) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\",\"0\",\"old\")")

	type NewDescription struct {
		Description string
	}
	var update NewDescription
	update.Description = "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/1/description", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/description", handler.UpdateDescription(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.Description != update.Description {
		t.Errorf("Expected to update password to \"%s\", but it is \"%s\"", update.Description, user.Description)
	}

}

// Updating a table with two users, updating the second user, correct id given
func TestUpdateDescription2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp,description) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\",\"0\",\"preserved\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp,description) values(\"DC\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\",\"0\",\"old\")")

	type NewDescription struct {
		Description string
	}
	var update NewDescription
	update.Description = "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/2/description", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/description", handler.UpdateDescription(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:True, ErrorExist:false} , but got %v", returnInfo)
	}

	var preservedUser handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser).First(&preservedUser, 1)
	globalDB.Model(&updatedUser).First(&updatedUser, 2)

	if preservedUser.Description != "preserved" {
		t.Errorf("Expected to preserve password as \"%s\", but it is \"%s\"", "preserved", preservedUser.Password)
	}

	if updatedUser.Description != update.Description {
		t.Errorf("Expected to update password to \"%s\", but it is \"%s\"", update.Description, updatedUser.Password)
	}

}

// Updating a table with two users, updating the third nonexsistent user
func TestUpdateDescription3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp,description) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\",\"0\",\"preserved\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password,xp,description) values(\"DC\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\",\"0\",\"preserved\")")

	type NewDescription struct {
		Description string
	}
	var update NewDescription
	update.Description = "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(update)

	req, err := http.NewRequest("PUT", "/api/users/3/description", &buf)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/description", handler.UpdateDescription(globalDB)).Methods("PUT")
	router.ServeHTTP(w, req)

	returnInfo := struct {
		ErrorExist bool
		Successful bool
	}{}

	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if returnInfo.Successful || !returnInfo.ErrorExist {
		t.Errorf("Expected {Successful:false, ErrorExist:True} , but got %v", returnInfo)
	}

	var preservedUser1 handler.User
	var preservedUser2 handler.User
	var updatedUser handler.User

	globalDB.Model(&preservedUser1).First(&preservedUser1, 1)
	globalDB.Model(&preservedUser2).First(&preservedUser2, 2)

	if preservedUser1.Description != "preserved" {
		t.Errorf("Expected to preserve description as \"%s\", but it is \"%s\"", "preserved", preservedUser1.Password)
	}

	if preservedUser2.Description != "preserved" {
		t.Errorf("Expected to preserve description as \"%s\", but it is \"%s\"", "preserved", preservedUser2.Password)
	}

	err = globalDB.Model(&handler.User{}).First(&updatedUser, 3).Error
	if err == nil { // There should not be a third user
		t.Errorf("Expected error, but user exsists: %v", updatedUser)
	}
}

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
	r, err := http.NewRequest("POST", "", &buf)
	if err != nil {
		panic(err)
	}
	handler.CheckLogin(globalDB)(w, r)

	returnInfo := struct {
		FindEmail    bool
		FindPassword bool
		User         handler.User
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.FindEmail || returnInfo.FindPassword {
		t.Errorf("Expected {FindEmail:true, FindPassword:false}, but got {%t %t}", returnInfo.FindEmail, returnInfo.FindPassword)
	}
}
