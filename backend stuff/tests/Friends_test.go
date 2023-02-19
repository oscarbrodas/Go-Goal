package test

import (
	"bytes"
	"encoding/json"
	"go-goal/httpd/handlers"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DSN string = "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

var globalDB *gorm.DB

func initializeTestDatabase() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("database failed to open")
	}
	globalDB = db

	// must drop goals and friends first before users because they have foreign keys in them
	globalDB.Exec("DROP TABLE goals")
	globalDB.Exec("DROP TABLE friends")
	globalDB.Exec("DROP TABLE users")

	globalDB.AutoMigrate(&handlers.User{})
	globalDB.AutoMigrate(&handlers.Goal{})
	globalDB.AutoMigrate(&handlers.Friend{})
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
	var user handlers.User
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
	handlers.GetAllFriends(globalDB)(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	expected := []int{2, 3}
	body := struct{ IDs []int }{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if !reflect.DeepEqual(body.IDs, expected) { // reflect.DeepEqual() is needed to compare slices
		t.Errorf("Expected %v, but got %v", expected, body.IDs)
	}
}
