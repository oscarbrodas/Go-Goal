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

func TestCreateGoal(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	inputGoal := handler.Goal{
		Title:       "my first goal",
		Description: "this is my first goal",
		Completed:   false,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputGoal)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/goals/1", &buf)

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/{id}", handler.CreateGoal(globalDB)).Methods("POST")
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	body := struct {
		Successful bool
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if body.ErrorExist { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("ErrorExist was true")
	}
	var exists bool
	globalDB.Model(&handler.Goal{}).Select("count(*) > 0").Where("user_id = 1").Find(&exists)
	if !exists {
		t.Errorf("did not find goal with user id 1")
	}
}

func TestGetGoals(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	globalDB.Exec("insert into goals(title, user_id) values(\"1\",1)")
	globalDB.Exec("insert into goals(title, user_id) values(\"2\",1)")
	globalDB.Exec("insert into goals(title, user_id) values(\"3\",1)")

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/goals/1", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/{id}", handler.GetGoals(globalDB)).Methods("GET")
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	body := struct {
		Successful bool
		ErrorExist bool
		Goals      []handler.Goal
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if body.ErrorExist { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("ErrorExist was true")
	}
	if body.Goals[0].Title != "1" || body.Goals[1].Title != "2" || body.Goals[2].Title != "3" {
		t.Errorf("Expected titles with 1 2 3 but got: %s %s %s", body.Goals[0].Title, body.Goals[1].Title, body.Goals[2].Title)
	}
}

// tests by deleteing goalID 2
func TestDeleteGoal1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	globalDB.Exec("insert into goals(id,title, user_id) values(1,\"1\",1)")
	globalDB.Exec("insert into goals(id,title, user_id) values(2,\"2\",1)")

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/goals/2", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/{goalID}", handler.DeleteGoal(globalDB)).Methods("DELETE")
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	body := struct {
		Successful bool
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if body.ErrorExist { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("ErrorExist was true")
	}
	var someGoal handler.Goal
	someGoal.Title = ""
	globalDB.First(&someGoal, 2)
	if someGoal.Title == "2" {
		t.Errorf("Found the tuple with title 2 but it is suppose to be deleted")
	}
}

// tests by deleting a goal that doesnt exist
func TestDeleteGoal2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	globalDB.Exec("insert into goals(id,title, user_id) values(1,\"1\",1)")

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/goals/3", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/{goalID}", handler.DeleteGoal(globalDB)).Methods("DELETE")
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	body := struct {
		Successful bool
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if body.ErrorExist { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("ErrorExist was true")
	}
}

func TestAddBenchmark(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	globalDB.Exec("insert into goals(id,title, user_id) values(1,\"1\",1)")

	bm := handler.Benchmark{
		Description: "first benchmark",
		Completed:   false,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(bm)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/goals/1/1", &buf)

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/{id}/{goalID}", handler.AddBenchmark(globalDB)).Methods("POST")
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	body := struct {
		Successful bool
		ErrorExist bool
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if body.ErrorExist { // reflect.DeepEqual() is needed to compare slices and structs
		t.Errorf("ErrorExist was true")
	}
	var exists bool
	globalDB.Model(&handler.Benchmark{}).Select("count(*) > 0").Where("goal_id = 1").Find(&exists)
	if !exists {
		t.Errorf("did not find benchmark with goal id 1")
	}
}

// tests to get no benchmarks
func TestGetBenchmarks1(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	globalDB.Exec("insert into goals(id,title, user_id) values(1,\"1\",1)")

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/goals/benchmarks/1", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/benchmarks/{goalID}", handler.GetBenchmarks(globalDB)).Methods("GET")
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	body := struct {
		Successful bool
		ErrorExist bool
		benches    []handler.Benchmark
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if body.ErrorExist {
		t.Errorf("ErrorExist was true")
	}
	if len(body.benches) != 0 {
		t.Errorf("Expected size of 0 but got size of %d", len(body.benches))
	}
}

// tests for when goalID does not exist
func TestGetBenchmarks2(t *testing.T) {
	initializeTestDatabase()

	globalDB.Exec("insert into users(first_name,last_name,email,password) values(\"1\",\"Chen\",\"1@gmail.com\",\"pw\")")

	globalDB.Exec("insert into goals(id,title, user_id) values(1,\"1\",1)")

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/goals/benchmarks/2", nil)

	router := mux.NewRouter()
	router.HandleFunc("/api/goals/benchmarks/{goalID}", handler.GetBenchmarks(globalDB)).Methods("GET")
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get StatusOK, instead got %d", w.Result().StatusCode)
	}

	body := struct {
		Successful bool
		ErrorExist bool
		benches    []handler.Benchmark
	}{}
	json.NewDecoder(w.Result().Body).Decode(&body)
	if !body.ErrorExist {
		t.Errorf("ErrorExist was false, expected true")
	}
	if len(body.benches) != 0 {
		t.Errorf("Expected size of 0 but got size of %d", len(body.benches))
	}
}
