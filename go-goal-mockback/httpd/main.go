package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-goal/httpd/handler"
	"go-goal/platform/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("../platform/user/user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&user.User{})

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/login", handler.CheckLogin(db)).Methods("GET")
	s.HandleFunc("/sign-up", handler.SignUp(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}
