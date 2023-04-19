package handler

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

func printError(thisUser uint, otherUser string, funcName string, message string) {
	fmt.Printf("Error in %s\nuserID:%d performed action on userID:%s\n", funcName, thisUser, otherUser)
	fmt.Print(message + "\n")
}

// http request must have the user object
// returns a json of an array of numbers of IDs
// the name of the array called "IDs"
func GetAllFriends(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userID uint64
		params := mux.Vars(r)
		userID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			panic(err)
		}
		var friends []Friend
		returnInfo := struct {
			IDs        []uint // name needs to be standardized
			ErrorExist bool
		}{}

		result := globalDB.Where("(user1 = ? OR user2 = ?) AND who_sent = 0", userID, userID).Find(&friends)
		if result.Error != nil {
			returnInfo.ErrorExist = true
			print(result.Error)
		} else {
			for i := 0; i < len(friends); i++ {
				if friends[i].User1 != uint(userID) {
					returnInfo.IDs = append(returnInfo.IDs, friends[i].User1)
				} else {
					returnInfo.IDs = append(returnInfo.IDs, friends[i].User2)
				}
			}
		}
		json.NewEncoder(w).Encode(returnInfo)
	}
}

// errorExist is true even if no error, but just invalid send friend request
func SendFriendRequest(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var senderID uint64
		var recieverID uint64
		params := mux.Vars(r)
		senderID, err := strconv.ParseUint(params["sender"], 10, 64)
		if err != nil {
			panic(err)
		}
		recieverID, err = strconv.ParseUint(params["reciever"], 10, 64)
		if err != nil {
			panic(err)
		}

		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
		}{}

		var exists1 bool
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ?", recieverID, senderID).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ?", senderID, recieverID).Find(&exists2)

		if exists1 || exists2 {
			returnInfo.ErrorExist = true
		} else {
			if err != nil {
				printError(uint(senderID), params["reciever"], "SendFriendRequest", "Input parameter in URL was not a number")
				returnInfo.ErrorExist = true
				json.NewEncoder(w).Encode(returnInfo)
				return
			}

			friendInput := Friend{User1: uint(senderID), User2: uint(recieverID), WhoSent: 1}
			result := globalDB.Model(&Friend{}).Create(&friendInput)
			if result.Error != nil {
				returnInfo.ErrorExist = true
				printError(uint(senderID), params["reciever"], "SendFriendRequest", "Error when inserting")
			} else {
				returnInfo.Successful = true
			}
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func GetOutgoingFriendRequests(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userID uint64
		params := mux.Vars(r)
		userID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			panic(err)
		}

		var friends []Friend
		returnInfo := struct { // need to be standardized
			IDs        []uint
			ErrorExist bool
		}{}

		globalDB.Where("user1 = ? AND who_sent = 1", userID).Find(&friends)
		for i := 0; i < len(friends); i++ {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User2)
		}

		globalDB.Where("user2 = ? AND who_sent = ?", userID, 2).Find(&friends)
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

		var userID uint64
		params := mux.Vars(r)
		userID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			panic(err)
		}

		var friends []Friend
		returnInfo := struct { // need to be standardized
			IDs        []uint
			ErrorExist bool
		}{}

		globalDB.Where("user1 = ? AND who_sent = ?", userID, 2).Find(&friends)
		for i := 0; i < len(friends); i++ {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User2)
		}

		globalDB.Where("user2 = ? AND who_sent = ?", userID, 1).Find(&friends)
		for i := 0; i < len(friends); i++ {
			returnInfo.IDs = append(returnInfo.IDs, friends[i].User1)
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func AcceptFriendRequest(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var senderID uint64
		var recieverID uint64
		params := mux.Vars(r)
		senderID, err := strconv.ParseUint(params["sender"], 10, 64)
		if err != nil {
			panic(err)
		}
		recieverID, err = strconv.ParseUint(params["accepter"], 10, 64)
		if err != nil {
			panic(err)
		}

		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
		}{}

		var exists1 bool //much better way to check if soemthing exists
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 1", senderID, recieverID).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 2", recieverID, senderID).Find(&exists2)

		if exists1 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 1", senderID, recieverID).Update("who_sent", 0)
		} else if exists2 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 2", recieverID, senderID).Update("who_sent", 0)
		} else {
			returnInfo.ErrorExist = true
			printError(uint(recieverID), params["sender"], "AcceptFriendRequest", "")
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func DeclineFriendRequest(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var senderID uint64
		var recieverID uint64
		params := mux.Vars(r)
		senderID, err := strconv.ParseUint(params["sender"], 10, 64)
		if err != nil {
			panic(err)
		}
		recieverID, err = strconv.ParseUint(params["decliner"], 10, 64)
		if err != nil {
			panic(err)
		}

		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
		}{}

		var exists1 bool //much better way to check if soemthing exists
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 1", senderID, recieverID).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 2", recieverID, senderID).Find(&exists2)

		if exists1 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 1", senderID, recieverID).Delete(&Friend{})
		} else if exists2 {
			returnInfo.Successful = true
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = 2", recieverID, senderID).Delete(&Friend{})
		} else {
			returnInfo.ErrorExist = true
			printError(uint(recieverID), params["sender"], "DeclineFriendRequest", "")
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
		var remover uint64
		var removedFriend uint64
		params := mux.Vars(r)
		remover, err := strconv.ParseUint(params["remover"], 10, 64)
		if err != nil {
			panic(err)
		}
		removedFriend, err = strconv.ParseUint(params["friend"], 10, 64)
		if err != nil {
			panic(err)
		}

		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
		}{}

		var exists1 bool //much better way to check if soemthing exists
		var exists2 bool
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 0", remover, removedFriend).Find(&exists1)
		globalDB.Model(&Friend{}).Select("count(*) > 0").Where("user1 = ? AND user2 = ? AND who_sent = 0", removedFriend, remover).Find(&exists2)

		if exists1 {
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = ?", remover, removedFriend, 0).Delete(&Friend{})
			returnInfo.Successful = true
		} else if exists2 {
			globalDB.Model(&Friend{}).Where("user1 = ? AND user2 = ? AND who_sent = ?", removedFriend, remover, 0).Delete(&Friend{})
			returnInfo.Successful = true
		} else {
			returnInfo.ErrorExist = true
			printError(uint(remover), params["friend"], "RemoveFriend", "")
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func SearchFriend(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		name, _ := params["name"]

		returnInfo := struct { // need to be standardized
			Successful bool
			ErrorExist bool
			users      []User
		}{}

		globalDB.Model(User{}).Where("username LIKE ? AND ROWNUM < 11", "%"+name+"%").Find(&returnInfo.users)

		json.NewEncoder(w).Encode(returnInfo)
	}
}
