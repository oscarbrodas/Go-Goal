package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-goal/httpd/handler"
	"go-goal/platform/user"
)

func main() {
	accounts := user.New()
	r := mux.NewRouter()
	//s := r.Host("/api").Subrouter() // Wasn't working, need to investigate

	r.HandleFunc("/api/login", handler.CheckLogin(accounts)).Methods("GET")
	r.HandleFunc("/api/sign-up", handler.SignUp(accounts)).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))
}
