package main

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model

	User1   int
	User2   int
	WhoSent int
	User    User `gorm:"foreignKey:User1"`
	Userx   User `gorm:"foreignKey:User2"`
}

// http request must have the ID of the user in json
// returns a json of an array of numbers of IDs
// the name of the array called "IDs"
func getAllFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friends []Friend
	var user User
	returnInfo := struct {
		IDs []int
	}{}
	json.NewDecoder(r.Body).Decode(&user)

	globalDB.Where("(user1 = ? OR user2 = ?) AND who_sent = ?", user.ID, user.ID, 0).Find(&friends)

	for i := 0; i < len(friends); i++ {
		if friends[i].User1 != int(user.ID) {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User1)
		} else {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User2)
		}
	}

	json.NewEncoder(w).Encode(returnInfo)
}
