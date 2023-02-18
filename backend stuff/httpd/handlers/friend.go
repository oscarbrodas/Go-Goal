package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model

	User1   uint
	User2   uint
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
			IDs []uint // name needs to be standardized
		}{}
		json.NewDecoder(r.Body).Decode(&user)

		globalDB.Where("(user1 = ? OR user2 = ?) AND who_sent = ?", user.ID, user.ID, 0).Find(&friends)

		for i := 0; i < len(friends); i++ {
			if friends[i].User1 != user.ID {
				returnInfo.IDs = append(returnInfo.IDs, friends[i].User1)
			} else {
				returnInfo.IDs = append(returnInfo.IDs, friends[i].User2)
			}
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input json must contain the user who is sending the request
// cannot send a request if you already sent one
// cannot send a request if the other person sent you one
// cannot send a request if you are already friends
// returns if the request failed or succeeded
func SendFriendRequest(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		otherID := params["id"]
		var thisUser User
		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
			CanSend    bool
		}{}
		json.NewDecoder(r.Body).Decode(&thisUser)

		var exists1 bool
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ?", otherID, thisUser.ID).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ?", thisUser.ID, otherID).Find(&exists2)

		if exists1 || exists2 {
			returnInfo.CanSend = false
		} else {
			otherIDInt, err := strconv.Atoi(otherID)
			if err != nil {
				str := strconv.FormatUint(uint64(thisUser.ID), 10)
				fmt.Print("\nError in SendFriendRequest\n" + str + " sent a request to userID:" + otherID + "\n")
				fmt.Print("value in the parameters of the route was not a number\n")
				returnInfo.ErrorExist = true
				json.NewEncoder(w).Encode(returnInfo)
				return
			}

			friendInput := Friend{User1: thisUser.ID, User2: uint(otherIDInt), WhoSent: 1}
			globalDB.Model(&Friend{}).Create(&friendInput)
			returnInfo.CanSend = true
			returnInfo.Successful = true
		}

		json.NewEncoder(w).Encode(returnInfo)
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
			IDs []uint
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
			IDs []uint
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

// input params must contain the ID of the user that sent the friend request
// input json must be of the user who accepted
// returns json of what happened
func AcceptFriendRequest(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		otherID := params["id"]
		var thisUser User
		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
		}{}
		json.NewDecoder(r.Body).Decode(&thisUser)

		var exists1 bool //much better way to check if soemthing exists
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 1", otherID, thisUser.ID).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 2", thisUser.ID, otherID).Find(&exists2)

		if exists1 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 1", otherID, thisUser.ID).Update("who_sent", 0)
		} else if exists2 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 2", thisUser.ID, otherID).Update("who_sent", 0)
		} else {
			returnInfo.ErrorExist = true
			str := strconv.FormatUint(uint64(thisUser.ID), 10)
			fmt.Print("\nError in DeclineFriendRequest\nuserID:" + str + " declined a request from userID:" + otherID + "\n")
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input params must contain the ID of the user that sent the friend request
// input json must be of the user who declined
// returns json of what happened
func DeclineFriendRequest(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		otherID := params["id"]
		var thisUser User
		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
		}{}
		json.NewDecoder(r.Body).Decode(&thisUser)

		var exists1 bool //much better way to check if soemthing exists
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 1", otherID, thisUser.ID).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 2", thisUser.ID, otherID).Find(&exists2)

		if exists1 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 1", otherID, thisUser.ID).Delete(&Friend{})
		} else if exists2 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 2", thisUser.ID, otherID).Delete(&Friend{})
		} else {
			returnInfo.ErrorExist = true
			str := strconv.FormatUint(uint64(thisUser.ID), 10)
			fmt.Print("\nError in DeclineFriendRequest\nuserID:" + str + " declined a request from userID:" + otherID + "\n")
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

// input json must contain the user who is removing
// input params must contain the other id of the other user
// returns json of what happened
func RemoveFriend(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		otherID := params["id"]
		var thisUser User
		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
		}{}
		json.NewDecoder(r.Body).Decode(&thisUser)

		var exists1 bool //much better way to check if soemthing exists
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 0", thisUser.ID, otherID).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 0", otherID, thisUser.ID).Find(&exists2)

		if exists1 {
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = ?", thisUser.ID, otherID, 0).Delete(&Friend{})
			returnInfo.Successful = true
		} else if exists2 {
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = ?", otherID, thisUser.ID, 0).Delete(&Friend{})
			returnInfo.Successful = true
		} else {
			returnInfo.ErrorExist = true
			str := strconv.FormatUint(uint64(thisUser.ID), 10)
			fmt.Print("\nError in RemoveFriend\n" + "userID:" + str + " removed userID:" + otherID + "\n")
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}
