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

func TestGetUser(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")

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

	if returnInfo.ThisUser.FirstName != "2" || returnInfo.ThisUser.Email != "2@gmail.com" || returnInfo.ThisUser.Password != "pw" {
		t.Errorf("Got the wrong user")
	}
}

// creating a new user without the email already existing
func TestCreateUser1(t *testing.T) {
	initializeTestDatabase()
	initializeTestFileSystem()

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
	handler.CreateUser(globalDB, globalUploader, globalDownloader)(w, r)

	returnInfo := struct {
		Successful bool
		ErrorExist bool
		EmailExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !returnInfo.Successful || returnInfo.ErrorExist || returnInfo.EmailExist {
		t.Errorf("Expected {Successful:true, ErrorExist:false, EmailExist:false}, but got %v", returnInfo)
	}

	var inputtedUser handler.User
	globalDB.Model(&handler.User{}).Raw("SELECT username, first_name, last_name, email, password FROM users WHERE id = ?", 1).Scan(&inputtedUser)
	if !reflect.DeepEqual(user, inputtedUser) { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("Expected %v, but got %v", user, inputtedUser)
	}
}

// creating a new user with the email existing
func TestCreateUser2(t *testing.T) {
	initializeTestDatabase()
	initializeTestFileSystem()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	var user handler.User
	user.Username = "dwan12345"
	user.FirstName = "don"
	user.LastName = "chen"
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
	handler.CreateUser(globalDB, globalUploader, globalDownloader)(w, r)

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
	handler.CreateUser(globalDB, globalUploader, globalDownloader)(w, r)

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

// Updating a table with a single user, correct id given
func TestUpdateUsername1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"old\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")

	updatedUsername := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedUsername)

	req, err := http.NewRequest("PUT", "/api/users/1/username", &buf)
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

	var user handler.User

	globalDB.Model(&user).First(&user, 1)

	if user.Username != updatedUsername {
		t.Errorf("Expected to update username to \"%s\", but it is \"%s\"", updatedUsername, user.Username)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateUsername2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"preserved\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"old\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\")")

	updatedUsername := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedUsername)

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

	if updatedUser.Username != updatedUsername {
		t.Errorf("Expected to update username to \"%s\", but it is \"%s\"", updatedUsername, updatedUser.Username)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateUsername3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"preserved\",\"Don\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"preserved\",\"Don\", \"Chen\",\"2@gmail.com\",\"pw\")")

	updatedUsername := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedUsername)

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
		t.Errorf("Expected error to exsist, user exsists: %v", updatedUser)
	}
}

// Updating a table with a single user, correct id given
func TestUpdateFirstname1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"old\",\"Chen\",\"1@gmail.com\",\"pw\")")

	updatedFirstname := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedFirstname)

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

	if user.FirstName != updatedFirstname {
		t.Errorf("Expected to update firstname to \"%s\", but it is \"%s\"", updatedFirstname, user.FirstName)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateFirstname2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"preserved\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"old\", \"Chen\",\"2@gmail.com\",\"pw\")")

	updatedFirstname := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedFirstname)

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

	if updatedUser.FirstName != updatedFirstname {
		t.Errorf("Expected to update firstname to \"%s\", but it is \"%s\"", updatedFirstname, updatedUser.FirstName)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateFirstname3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"preserved\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"preserved\", \"Chen\",\"2@gmail.com\",\"pw\")")

	updatedFirstname := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedFirstname)

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

	updatedLastname := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedLastname)

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

	if user.LastName != updatedLastname {
		t.Errorf("Expected to update lastname to \"%s\", but it is \"%s\"", updatedLastname, user.LastName)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateLastname2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"preserved\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"old\",\"2@gmail.com\",\"pw\")")

	updatedLastname := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedLastname)

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

	if updatedUser.LastName != updatedLastname {
		t.Errorf("Expected to update lastname to \"%s\", but it is \"%s\"", updatedLastname, updatedUser.LastName)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateLastname3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"preserved\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"preserved\",\"2@gmail.com\",\"pw\")")

	updatedLastname := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedLastname)

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

	updatedEmail := "new@gmail.com"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedEmail)

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

	if user.Email != updatedEmail {
		t.Errorf("Expected to update email to \"%s\", but it is \"%s\"", updatedEmail, user.Email)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdateEmail2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"preserved@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"old@gmail.com\",\"pw\")")

	updatedEmail := "new@gmail.com"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedEmail)

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

	if updatedUser.Email != updatedEmail {
		t.Errorf("Expected to update eamil to \"%s\", but it is \"%s\"", updatedEmail, updatedUser.Email)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdateEmail3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"preserved1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"preserved2@gmail.com\",\"pw\")")

	updatedEmail := "new@gmail.com"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedEmail)

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

// Updating a table with a single user, correct id given
func TestUpdatePassword1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"old\")")

	updatedPassword := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedPassword)

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

	if user.Password != updatedPassword {
		t.Errorf("Expected to update password to \"%s\", but it is \"%s\"", updatedPassword, user.Password)
	}

}

// Updating a table with a two users, updating the second user, correct id given
func TestUpdatePassword2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"preserved\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"2@gmail.com\",\"old\")")

	updatedPassword := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedPassword)

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

	if updatedUser.Password != updatedPassword {
		t.Errorf("Expected to update password to \"%s\", but it is \"%s\"", updatedPassword, updatedUser.Password)
	}

}

// Updating a table with a two users, updating the third nonexsistent user
func TestUpdatePassword3(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\",\"Chen\",\"1@gmail.com\",\"preserved\")")
	globalDB.Exec("insert into users(username,first_name,last_name,email,password) values(\"DC\",\"Don\", \"Chen\",\"2@gmail.com\",\"preserved\")")

	updatedPassword := "new"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updatedPassword)

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
