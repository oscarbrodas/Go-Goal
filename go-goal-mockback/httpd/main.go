package main

import (
	"log"
	"os"
	"net/http"

	"go-goal/httpd/handler"
	"go-goal/platform/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := gorm.Open(sqlite.Open("../platform/user/user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&user.User{})

	
	host := "127.0.0.1:5000"
	if err := http.ListenAndServe(host, httpHandler(db)); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}


func httpHandler(db *gorm.DB) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/login", handler.CheckLogin(db)).Methods("GET")
	s.HandleFunc("/sign-up", handler.SignUp(db)).Methods("POST")


	r.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:5000", "http://localhost:4200"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(r))
}