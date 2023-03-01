package handler_test

import (
	"bytes"
	"encoding/json"
	"go-goal/httpd/handler"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func initializeTestDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("database failed to open")
	}
	globalDB = db

	// must drop goals and friends first before users because they have foreign keys in them
	globalDB.Exec("DROP TABLE goals")
	globalDB.Exec("DROP TABLE friends")
	globalDB.Exec("DROP TABLE users")

	globalDB.AutoMigrate(&handler.User{})
	globalDB.AutoMigrate(&handler.Goal{})
	globalDB.AutoMigrate(&handler.Friend{})
}

func TestGetAllFriends(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"3\",\"Chen\",\"3@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"4\",\"Chen\",\"4@gmail.com\",\"pw\")")

	globalDB.Exec("insert into friends(user1,user2,who_sent) values(1,2,0)")
	globalDB.Exec("insert into friends(user1,user2,who_sent) values(3,1,0)")
	globalDB.Exec("insert into friends(user1,user2,who_sent) values(4,1,1)")

	// all of this just sets the input json to a user with id:1
	var user handler.User
	user.ID = 1
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}

	// w will be a replacement for w http.ResponseWriter
	w := httptest.NewRecorder()

	// r will be a a replacement for r *http.Request
	// second parameter is for inputing stuff with the links such as /{id}
	// third parameter is the input json
	r, err := http.NewRequest(http.MethodGet, "", &buf)
	if err != nil {
		panic(err)
	}

	// this looks weird because GetAllFriends returns a handler function
	handler.GetAllFriends(globalDB)(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	expected := []int{2, 3}
	body := struct {
		IDs        []int
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if !reflect.DeepEqual(body.IDs, expected) { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("Expected %v, but got %v", expected, body.IDs)
	}
	if body.ErrorExist {
		t.Errorf("There was an error")
	}
}

func TestSendFriendRequest1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"3\",\"Chen\",\"3@gmail.com\",\"pw\")")

	var user handler.User
	user.ID = 1
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/?id=3", &buf)
	if err != nil {
		panic(err)
	}

	handler.SendFriendRequest(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}
	expected := struct {
		Successful bool
		ErrorExist bool
	}{
		Successful: true,
		ErrorExist: false,
	}
	returnInfo := struct {
		Successful bool
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !reflect.DeepEqual(returnInfo, expected) { // reflect.DeepEqual() is needed to compare slices
		t.Errorf("Expected %+v, but got %+v", expected, returnInfo)
	}
	var exists bool
	globalDB.Model(&handler.Friend{}).Select("count(*) > 0").Where("(user1 = 1 AND user2 = 3) AND who_sent = 1").Find(&exists)
	if !exists {
		t.Errorf("Did not find inserted tuple when 1 sent a friend request to 3")
	}
}

func TestSendFriendRequest2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"3\",\"Chen\",\"3@gmail.com\",\"pw\")")

	globalDB.Exec("insert into friends(user1,user2,who_sent) values(1,2,0)")

	var user handler.User
	user.ID = 1
	expected := struct {
		Successful bool
		ErrorExist bool
	}{
		Successful: false,
		ErrorExist: true,
	}
	returnInfo := struct {
		Successful bool
		ErrorExist bool
	}{}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/?id=2", &buf)
	if err != nil {
		panic(err)
	}

	handler.SendFriendRequest(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !reflect.DeepEqual(returnInfo, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, returnInfo)
	}
}

func TestGetOutgoingFriendRequests(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"3\",\"Chen\",\"3@gmail.com\",\"pw\")")

	globalDB.Exec("insert into friends(user1,user2,who_sent) values(1,2,1)")
	globalDB.Exec("insert into friends(user1,user2,who_sent) values(3,1,2)")

	var user handler.User
	user.ID = 1
	expected := struct {
		IDs        []uint
		ErrorExist bool
	}{
		IDs:        []uint{2, 3},
		ErrorExist: false,
	}
	returnInfo := struct {
		IDs        []uint
		ErrorExist bool
	}{}

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

	handler.GetOutgoingFriendRequests(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !reflect.DeepEqual(returnInfo, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, returnInfo)
	}
}

func TestGetIngoingFriendRequests(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"3\",\"Chen\",\"3@gmail.com\",\"pw\")")

	globalDB.Exec("insert into friends(user1,user2,who_sent) values(1,2,1)")
	globalDB.Exec("insert into friends(user1,user2,who_sent) values(3,2,1)")

	var user handler.User
	user.ID = 2
	expected := struct {
		IDs        []uint
		ErrorExist bool
	}{
		IDs:        []uint{1, 3},
		ErrorExist: false,
	}
	returnInfo := struct {
		IDs        []uint
		ErrorExist bool
	}{}

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

	handler.GetIngoingFriendRequests(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !reflect.DeepEqual(returnInfo, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, returnInfo)
	}
}

func TestAcceptFriendRequest(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")

	globalDB.Exec("insert into friends(user1,user2,who_sent) values(1,2,1)")

	var user handler.User
	user.ID = 2
	expected := struct {
		Successful bool
		ErrorExist bool
	}{
		Successful: true,
		ErrorExist: false,
	}
	returnInfo := struct {
		Successful bool
		ErrorExist bool
	}{}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("PUT", "/?id=1", &buf)
	if err != nil {
		panic(err)
	}

	handler.AcceptFriendRequest(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !reflect.DeepEqual(returnInfo, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, returnInfo)
	}
	var exists bool
	globalDB.Model(&handler.Friend{}).Select("count(*) > 0").Where("(user1 = 1 AND user2 = 2) AND who_sent = 0").Find(&exists)
	if !exists {
		t.Errorf("Did not find the tuple {user1 = 1, user2 = 2, who_sent = 0}")
	}
}

func TestDeclineFriendRequest(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")

	globalDB.Exec("insert into friends(user1,user2,who_sent) values(1,2,1)")

	var user handler.User
	user.ID = 2
	expected := struct {
		Successful bool
		ErrorExist bool
	}{
		Successful: true,
		ErrorExist: false,
	}
	returnInfo := struct {
		Successful bool
		ErrorExist bool
	}{}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/?id=1", &buf)
	if err != nil {
		panic(err)
	}

	handler.DeclineFriendRequest(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !reflect.DeepEqual(returnInfo, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, returnInfo)
	}
	var exists bool
	globalDB.Model(&handler.Friend{}).Select("count(*) > 0").Where("(user1 = 1 AND user2 = 2) AND who_sent = 1").Find(&exists)
	if exists {
		t.Errorf("Found the tuple {user1 = 1, user2 = 2, who_sent = 1} even though it should be deleted")
	}
}

func TestRemoveFriend(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")
	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"2\",\"Chen\",\"2@gmail.com\",\"pw\")")

	globalDB.Exec("insert into friends(user1,user2,who_sent) values(1,2,0)")

	var user handler.User
	user.ID = 2
	expected := struct {
		Successful bool
		ErrorExist bool
	}{
		Successful: true,
		ErrorExist: false,
	}
	returnInfo := struct {
		Successful bool
		ErrorExist bool
	}{}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/?id=1", &buf)
	if err != nil {
		panic(err)
	}

	handler.RemoveFriend(globalDB)(w, r)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}
	json.NewDecoder(w.Result().Body).Decode(&returnInfo)
	if !reflect.DeepEqual(returnInfo, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, returnInfo)
	}
	var exists bool
	globalDB.Model(&handler.Friend{}).Select("count(*) > 0").Where("(user1 = 1 AND user2 = 2) AND who_sent = 0").Find(&exists)
	if exists {
		t.Errorf("Found the tuple {user1 = 1, user2 = 2, who_sent = 0} even though it should be deleted")
	}
}
