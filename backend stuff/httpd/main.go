package main

import (
	"go-goal/httpd/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

const DSN string = "root:password@tcp(127.0.0.1:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"

// example route: http://localhost:9000/users

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers(globalDB)).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser(globalDB)).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser(globalDB)).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser(globalDB)).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser(globalDB)).Methods("DELETE")
	r.HandleFunc("/login", handlers.CheckLogin(globalDB)).Methods("GET")

	r.HandleFunc("/goals/{userID}", handlers.CreateGoal(globalDB)).Methods("POST")
	r.HandleFunc("/goals/{userID}", handlers.GetGoals(globalDB)).Methods("GET")

	r.HandleFunc("/friends", handlers.GetAllFriends(globalDB)).Methods("GET")
	r.HandleFunc("/friends/sendFriendRequest", handlers.SendFriendRequest(globalDB)).Methods("POST") // the route should be changed
	r.HandleFunc("/friends/getOutgoingFriendRequests", handlers.GetOutgoingFriendRequests(globalDB)).Methods("GET")
	r.HandleFunc("/friends/getIngoingFriendRequests", handlers.GetIngoingFriendRequests(globalDB)).Methods("GET")
	r.HandleFunc("/friends/acceptFriendRequest", handlers.AcceptFriendRequest(globalDB)).Methods("PUT")
	r.HandleFunc("/friends/declineFriendRequest", handlers.DeclineFriendRequest(globalDB)).Methods("DELETE")
	r.HandleFunc("/friends/removeFriend", handlers.RemoveFriend(globalDB)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r)) // :9000 is the port
}

func main() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("database failed to open")
	}
	globalDB = db

	// MAKE SURE TO AUTOMIGRATE BEFORE INITALIZEROUTER
	globalDB.AutoMigrate(&handlers.User{})
	globalDB.AutoMigrate(&handlers.Goal{})
	globalDB.AutoMigrate(&handlers.Friend{})

	initializeRouter()
}
