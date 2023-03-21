package handler

import (
	"encoding/json"
	"fmt"
	"go-goal/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model

	Title       string
	Description string
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}

// input body contain "ThisUser" and "ThisGoal" objects
// "ThisGoal" does not need to contain UserID attribute
func CreateGoal(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var InputGoal Goal
		var UserID uint64
		params := mux.Vars(r)
		UserID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			panic(err)
		}
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		util.DecodeJSONRequest(&InputGoal, r.Body, w)
		InputGoal.UserID = uint(UserID)

		var userExists bool
		globalDB.Model(&User{}).Select("count(*) > 0").Where("id = ?", UserID).Find(&userExists)
		if userExists {
			result := globalDB.Model(&Goal{}).Create(&InputGoal)
			if result.Error != nil {
				returnInfo.ErrorExist = true
				fmt.Println(result.Error)
			} else {
				returnInfo.Successful = true
			}
		} else {
			returnInfo.ErrorExist = true
			fmt.Println("Error in Create Goal:")
			fmt.Printf("User:%d does not exist\n", UserID)
		}
		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input body contain user object
func GetGoals(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var UserID uint64
		params := mux.Vars(r)
		UserID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			panic(err)
		}
		returnInfo := struct {
			Successful bool
			ErrorExist bool
			Goals      []Goal
		}{}

		var userExists bool
		globalDB.Model(&User{}).Select("count(*) > 0").Where("id = ?", UserID).Find(&userExists)

		if userExists {
			globalDB.Model(&Goal{}).Where("user_id = ?", UserID).Find(&returnInfo.Goals)
			returnInfo.Successful = true
		} else {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in GetGoals\nUser:%d does not exist\n", UserID)
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input body contain the goal object to delete
func DeleteGoal(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var goalID uint64
		params := mux.Vars(r)
		goalID, err := strconv.ParseUint(params["goalID"], 10, 64)
		if err != nil {
			panic(err)
		}
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}

		result := globalDB.Delete(&Goal{}, goalID)
		if result.Error != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in DeleteGoal with deleting GoalID:%d\n", goalID)
		} else {
			returnInfo.Successful = true
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}
