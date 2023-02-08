package main

import (
	"log"
	"net/http"
	"go-goal/httpd/handlers"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

const DSN string = "root:password@tcp(127.0.0.1:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers(globalDB)).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser(globalDB)).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser(globalDB)).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser(globalDB)).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser(globalDB)).Methods("DELETE")
	r.HandleFunc("/login", handlers.CheckLogin(globalDB)).Methods("GET")

	r.HandleFunc("/goals/{userID}", handlers.CreateGoal(globalDB)).Methods("POST")
	r.HandleFunc("/goals/{userID}", handlers.GetGoals(globalDB)).Methods("GET")

	r.HandleFunc("/friends", handlers.GetAllFriends(globalDB)).Methods("GET")
	r.HandleFunc("/friends/sendFriendRequest", handlers,SendFriendRequest(globalDB)).Methods("POST") // the route should be changed
	r.HandleFunc("/friends/getOutgoingFriendRequests", handlers.GetOutgoingFriendRequests(globalDB)).Methods("GET")
	r.HandleFunc("/friends/getIngoingFriendRequests", handlers.GetIngoingFriendRequests(globalDB)).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r)) // :9000 is the port
}

func main() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("database failed to open")
	}
	globalDB = db

	// MAKE SURE TO AUTOMIGRATE BEFORE INITALIZEROUTER
	globalDB.AutoMigrate(&User{})
	globalDB.AutoMigrate(&Goal{})
	globalDB.AutoMigrate(&Friend{})

	initializeRouter()

}
