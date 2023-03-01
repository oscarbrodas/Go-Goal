package handler_test

import (
	"bytes"
	"encoding/json"
	"go-goal/httpd/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Test if a goal is succesfully inserted into the database
func TestCreateGoal(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	inputUser := handler.User{}

	inputGoal := handler.Goal{
		Title:       "Title",
		Description: "Description",
		UserID:      1,
		User:        inputUser,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputGoal)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/goals/1", &buf)
	if err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/{userID}", handler.CreateGoal(globalDB)).Methods("POST")
	router.ServeHTTP(w, req)

	var returnGoal handler.Goal

	json.NewDecoder(w.Result().Body).Decode(&returnGoal)
	if (returnGoal.Title != inputGoal.Title) && (returnGoal.Description != inputGoal.Description) && (returnGoal.UserID != inputGoal.UserID) && (returnGoal.User != inputGoal.User) {
		t.Errorf("Expected passing in goal to be return. Input Goal: %v Return Goal: %v", inputGoal, returnGoal)
	}

	var tableGoal handler.Goal

	globalDB.Model(&tableGoal).Where("user_id = ?", 1).Find(&tableGoal)

	if (tableGoal.Title != inputGoal.Title) && (tableGoal.Description != inputGoal.Description) && (tableGoal.UserID != inputGoal.UserID) && (tableGoal.User != inputGoal.User) {
		t.Errorf("Expected passing in goal to be inputed into the table. Input Goal: %v Table Goal: %v", inputGoal, tableGoal)
	}

}
