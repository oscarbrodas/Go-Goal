package handlers

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

// http request must have the user object
// returns a json of an array of numbers of IDs
// the name of the array called "IDs"
func GetAllFriends(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var friends []Friend
		var user User
		returnInfo := struct {
			IDs []int // name needs to be standardized
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
}

// input json must contain "user1" and "user2". user1 sends a request to user2
// cannot send a request if you already sent one
// cannot send a request if the other person sent you one
// cannot send a request if you are already friends
// returns if the request failed or succeeded
func SendFriendRequest(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		inputUsers := struct {
			User1 User
			User2 User
		}{}
		var canSend bool
		var c1 int64
		var c2 int64

		json.NewDecoder(r.Body).Decode(&inputUsers)

		globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ?", inputUsers.User1.ID, inputUsers.User2.ID).Count(&c1)
		globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ?", inputUsers.User2.ID, inputUsers.User1.ID).Count(&c2)
		canSend = (c1 == 0) && (c2 == 0)

		if canSend {
			var friendInput Friend
			friendInput.User1 = int(inputUsers.User1.ID)
			friendInput.User2 = int(inputUsers.User2.ID)
			friendInput.WhoSent = 1
			globalDB.Model(&Friend{}).Create(&friendInput)

			json.NewEncoder(w).Encode(struct{ RequestSent bool }{RequestSent: true})
		} else {
			json.NewEncoder(w).Encode(struct{ RequestSent bool }{RequestSent: false})
		}
	}
}

// the input json must be of the user object
// returns a json of an array of the user IDs named "IDs"
func GetOutgoingFriendRequests(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user User
		var friends []Friend
		returnInfo := struct { // need to be standardized
			IDs []int
		}{}

		json.NewDecoder(r.Body).Decode(&user)
		globalDB.Where("user1 = ? AND who_sent = ?", user.ID, 1).Find(&friends)
		for i := 0; i < len(friends); i++ {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User2)
		}

		json.NewDecoder(r.Body).Decode(&user)
		globalDB.Where("user2 = ? AND who_sent = ?", user.ID, 2).Find(&friends)
		for i := 0; i < len(friends); i++ {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User1)
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input json must contain the user object
// return will contain json of an array called "IDs" that contains IDs
func GetIngoingFriendRequests(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user User
		var friends []Friend
		returnInfo := struct { // need to be standardized
			IDs []int
		}{}

		json.NewDecoder(r.Body).Decode(&user)
		globalDB.Where("user1 = ? AND who_sent = ?", user.ID, 2).Find(&friends)
		for i := 0; i < len(friends); i++ {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User2)
		}

		json.NewDecoder(r.Body).Decode(&user)
		globalDB.Where("user2 = ? AND who_sent = ?", user.ID, 1).Find(&friends)
		for i := 0; i < len(friends); i++ {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User1)
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}
