package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model

	Title       string
	Description string
	UserID      int
	User        User `gorm:"foreignKey:UserID"`
}

func CreateGoal(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var goal Goal
		var err error
		params := mux.Vars(r)

		json.NewDecoder(r.Body).Decode(&goal)
		goal.UserID, err = strconv.Atoi(params["userID"])
		if err != nil {
			panic("userID of goal POST request was not an int")
		}
		globalDB.Create(&goal)
		json.NewEncoder(w).Encode(goal)
	}
}

func GetGoals(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var goal []Goal
		params := mux.Vars(r)

		globalDB.Where("user_id = ?", params["userID"]).Find(&goal)

		json.NewEncoder(w).Encode(goal)
	}
}
