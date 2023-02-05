package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

const DSN string = "root:password@tcp(127.0.0.1:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/login", checkLogin).Methods("GET")

	r.HandleFunc("/goals/{userID}", createGoal).Methods("POST")
	r.HandleFunc("/goals/{userID}", getGoals).Methods("GET")

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

	initializeRouter()

}
