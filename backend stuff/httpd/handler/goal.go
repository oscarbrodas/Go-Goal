package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		input := struct {
			ThisGoal Goal
			ThisUser User
		}{}
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		json.NewDecoder(r.Body).Decode(&input)
		input.ThisGoal.UserID = uint(input.ThisUser.ID)

		var userExists bool
		globalDB.Model(&User{}).Select("count(*) > 0").Where("id = ?", input.ThisGoal.UserID).Find(&userExists)
		if userExists {
			result := globalDB.Model(&Goal{}).Create(&input.ThisGoal)
			if result.Error != nil {
				returnInfo.ErrorExist = true
				fmt.Println(result.Error)
			} else {
				returnInfo.Successful = true
			}
		} else {
			returnInfo.ErrorExist = true
			fmt.Println("Error in Create Goal:")
			fmt.Printf("User:%d does not exist\n", input.ThisGoal.UserID)
		}
		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input body contain user object
func GetGoals(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var ThisUser User
		returnInfo := struct {
			Successful bool
			ErrorExist bool
			Goals      []Goal
		}{}
		json.NewDecoder(r.Body).Decode(&ThisUser)

		var userExists bool
		globalDB.Model(&User{}).Select("count(*) > 0").Where("id = ?", ThisUser.ID).Find(&userExists)

		if userExists {
			globalDB.Model(&Goal{}).Where("user_id = ?", ThisUser.ID).Find(&returnInfo.Goals)
			returnInfo.Successful = true
		} else {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in GetGoals\nUser:%d does not exist\n", ThisUser.ID)
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input body contain the goal object to delete
func DeleteGoal(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var ThisGoal Goal
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		json.NewDecoder(r.Body).Decode(&ThisGoal)

		result := globalDB.Delete(&ThisGoal)
		if result.Error != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in DeleteGoal with deleting GoalID:%d\n", ThisGoal.ID)
		} else {
			returnInfo.Successful = true
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}
