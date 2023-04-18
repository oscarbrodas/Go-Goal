package handler_test

import (
	"context"
	"encoding/json"
	"go-goal/httpd/handler"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var globalDB *gorm.DB
var globalUploader *manager.Uploader
var globalDownloader *manager.Downloader

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

func initializeTestFileSystem() {

	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup s3 uploader
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	globalUploader = manager.NewUploader(client)
	globalDownloader = manager.NewDownloader(client)
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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/api/friends/1", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/{id}", handler.GetAllFriends(globalDB)).Methods("GET")
	router.ServeHTTP(w, r)

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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/friends/sendFriendRequest/1/3", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/sendFriendRequest/{sender}/{reciever}", handler.SendFriendRequest(globalDB)).Methods("POST")
	router.ServeHTTP(w, r)

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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/friends/sendFriendRequest/1/2", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/sendFriendRequest/{sender}/{reciever}", handler.SendFriendRequest(globalDB)).Methods("POST")
	router.ServeHTTP(w, r)

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
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/api/friends/getOutgoingFriendRequests/1", nil)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/getOutgoingFriendRequests/{id}", handler.GetOutgoingFriendRequests(globalDB)).Methods("GET")
	router.ServeHTTP(w, r)

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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/friends/getIngoingFriendRequests/2", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/getIngoingFriendRequests/{id}", handler.GetIngoingFriendRequests(globalDB)).Methods("GET")
	router.ServeHTTP(w, r)

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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/friends/acceptFriendRequest/1/2", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/acceptFriendRequest/{sender}/{accepter}", handler.AcceptFriendRequest(globalDB)).Methods("PUT")
	router.ServeHTTP(w, r)

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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/friends/declineFriendRequest/1/2", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/declineFriendRequest/{sender}/{decliner}", handler.DeclineFriendRequest(globalDB)).Methods("DELETE")
	router.ServeHTTP(w, r)

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

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/friends/removeFriend/1/2", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/friends/removeFriend/{remover}/{friend}", handler.RemoveFriend(globalDB)).Methods("DELETE")
	router.ServeHTTP(w, r)

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
