package main

// Parts Taken From Dobra's Code

import (
	"net/http"
	"os"

	"go-goal/httpd/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func httpHandler() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter() // adds "/api" to each route

	/* ====================== REST API ====================== */

	// Create, Update, Retrieve, Handle Users
	s.HandleFunc("/users", handler.GetUser(globalDB)).Methods("GET")
	s.HandleFunc("/users", handler.CreateUser(globalDB)).Methods("POST")
	s.HandleFunc("/users/{id}/username", handler.UpdateUsername(globalDB)).Methods("PUT")
	s.HandleFunc("/users/{id}/firstname", handler.UpdateFirstname(globalDB)).Methods("PUT")
	s.HandleFunc("/users/{id}/lastname", handler.UpdateLastname(globalDB)).Methods("PUT")
	s.HandleFunc("/users/{id}/email", handler.UpdateEmail(globalDB)).Methods("PUT")
	s.HandleFunc("/users/{id}/password", handler.UpdatePassword(globalDB)).Methods("PUT")
	s.HandleFunc("/users/{id}/xp", handler.AddXP(globalDB)).Methods("PUT")
	s.HandleFunc("/users/{id}/description", handler.UpdateDescription(globalDB)).Methods("PUT")
	s.HandleFunc("/login/{email}/{password}", handler.CheckLogin(globalDB)).Methods("GET")
	s.HandleFunc("/users/checkUsername/{username}", handler.CheckUsername(globalDB)).Methods("GET")
	s.HandleFunc("/users/{id}/avatar", handler.SetAvatar(globalDB, globalUploader)).Methods("PUT")
	s.HandleFunc("/users/{id}/avatar", handler.GetAvatar(globalDB, globalDownloader)).Methods("GET")

	// Create and Retrieve Goals
	s.HandleFunc("/goals/{id}", handler.CreateGoal(globalDB)).Methods("POST")
	s.HandleFunc("/goals/{id}", handler.GetGoals(globalDB)).Methods("GET")
	s.HandleFunc("/goals/{goalID}", handler.DeleteGoal(globalDB)).Methods("DELETE")
	s.HandleFunc("/goals/{goalID}", handler.UpdateGoal(globalDB)).Methods("PUT")
	s.HandleFunc("/goals/{id}/{goalID}", handler.AddBenchmark(globalDB)).Methods("POST")
	s.HandleFunc("/goals/benchmarks/{goalID}", handler.GetBenchmarks(globalDB)).Methods("GET")
	s.HandleFunc("/goals/benchmarks/description/{benchmarkID}", handler.UpdateBenchmarkDescription(globalDB)).Methods("PUT")
	s.HandleFunc("/goals/benchmarks/completion/{benchmarkID}", handler.UpdateBenchmarkCompletion(globalDB)).Methods("PUT")
	s.HandleFunc("/goals/benchmarks/{benchmarkID}", handler.DeleteBenchmark(globalDB)).Methods("DELETE")

	// Retrieve/Remove Friends, Handle Friend Requests
	s.HandleFunc("/friends/{id}", handler.GetAllFriends(globalDB)).Methods("GET")
	s.HandleFunc("/friends/sendFriendRequest/{sender}/{reciever}", handler.SendFriendRequest(globalDB)).Methods("POST") // the route should be changed
	s.HandleFunc("/friends/getOutgoingFriendRequests/{id}", handler.GetOutgoingFriendRequests(globalDB)).Methods("GET")
	s.HandleFunc("/friends/getIngoingFriendRequests/{id}", handler.GetIngoingFriendRequests(globalDB)).Methods("GET")
	s.HandleFunc("/friends/acceptFriendRequest/{sender}/{accepter}", handler.AcceptFriendRequest(globalDB)).Methods("PUT")
	s.HandleFunc("/friends/declineFriendRequest/{sender}/{decliner}", handler.DeclineFriendRequest(globalDB)).Methods("DELETE")
	s.HandleFunc("/friends/removeFriend/{remover}/{friend}", handler.RemoveFriend(globalDB)).Methods("DELETE")

	// Route to serve site - MUST BE FINAL ROUTE
	r.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	// CORRS Handling, Imported by gorilla/handlers
	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:9000", "http://localhost:4200"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(r))
}
