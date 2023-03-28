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
	Completed   bool
	User        User `gorm:"foreignKey:UserID"`
}

type Benchmark struct {
	gorm.Model

	Description string
	Completed   bool
	GoalID      uint
	Goal        Goal `gorm:"foreignKey:ID"`
}

// input body contain "ThisUser" and "ThisGoal" objects
// "ThisGoal" does not need to contain UserID attribute
func CreateGoal(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var InputGoal Goal
		var UserID uint64
		params := mux.Vars(r)
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		UserID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}

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
		returnInfo := struct {
			Successful bool
			ErrorExist bool
			Goals      []Goal
		}{}
		UserID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}

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
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		goalID, err := strconv.ParseUint(params["goalID"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}

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

// input body contain the goal object to update to
func UpdateGoal(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var goalID uint64
		params := mux.Vars(r)
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		goalID, err := strconv.ParseUint(params["goalID"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}

		var newGoal Goal
		util.DecodeJSONRequest(&newGoal, r.Body, w)

		result := globalDB.Model(&newGoal).Where("id = ?", goalID).Select("Title", "Description", "Completed").Updates(
			Goal{Title: newGoal.Title, Description: newGoal.Description, Completed: newGoal.Completed})
		if result.Error != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in UpdateGoal with updating GoalID:%d\n", goalID)
		} else {
			returnInfo.Successful = true
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func AddBenchmark(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var InputBenchmark Benchmark
		var UserID uint64
		var GoalID uint64
		params := mux.Vars(r)
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		UserID, err := strconv.ParseUint(params["id"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}
		GoalID, err = strconv.ParseUint(params["goalID"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}

		util.DecodeJSONRequest(&InputBenchmark, r.Body, w)
		InputBenchmark.GoalID = uint(GoalID)

		var userExists bool
		globalDB.Model(&User{}).Select("count(*) > 0").Where("id = ?", UserID).Find(&userExists)
		var goalExists bool
		globalDB.Model(&Goal{}).Select("count(*) > 0").Where("id = ?", GoalID).Find(&goalExists)
		if userExists && goalExists {
			result := globalDB.Model(&Benchmark{}).Create(&InputBenchmark)
			if result.Error != nil {
				returnInfo.ErrorExist = true
				fmt.Println(result.Error)
			} else {
				returnInfo.Successful = true
			}
		} else {
			returnInfo.ErrorExist = true
			fmt.Println("Error in AddBenchmark:")
			fmt.Printf("User:%d or Goal:%d does not exist\n", UserID, GoalID)
		}
		json.NewEncoder(w).Encode(returnInfo)
	}
}

func GetBenchmarks(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var GoalID uint64
		params := mux.Vars(r)
		returnInfo := struct {
			Successful bool
			ErrorExist bool
			Benchmarks []Benchmark
		}{}
		GoalID, err := strconv.ParseUint(params["goalID"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}

		var goalExists bool
		globalDB.Model(&Goal{}).Select("count(*) > 0").Where("id = ?", GoalID).Find(&goalExists)

		if goalExists {
			globalDB.Model(&Benchmark{}).Where("goal_id = ?", GoalID).Find(&returnInfo.Benchmarks)
			returnInfo.Successful = true
		} else {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in GetBenchmarks\nGoal:%d does not exist\n", GoalID)
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateBenchmarkDescription(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		returnInfo := struct {
			ErrorExist bool
			Successful bool
		}{}
		benchmarkID, err := strconv.ParseUint(params["benchmarkID"], 10, 64)
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Println(err)
		}

		var InputBenchmark Benchmark
		util.DecodeJSONRequest(&InputBenchmark, r.Body, w)
		InputBenchmark.ID = uint(benchmarkID)

		result := globalDB.Model(&Benchmark{}).Where("id = ?", benchmarkID).Update("description", InputBenchmark.Description)
		if result.Error != nil {
			returnInfo.ErrorExist = true
		} else {
			returnInfo.Successful = true
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func UpdateBenchmarkCompletion(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		returnInfo := struct {
			ErrorExist bool
			Successful bool
		}{}
		benchmarkID, err := strconv.ParseUint(params["benchmarkID"], 10, 64)
		if err != nil {
			returnInfo.ErrorExist = true
			fmt.Println(err)
		}

		var InputBenchmark Benchmark
		util.DecodeJSONRequest(&InputBenchmark, r.Body, w)
		InputBenchmark.ID = uint(benchmarkID)

		result := globalDB.Model(&Benchmark{}).Where("id = ?", benchmarkID).Update("completed", InputBenchmark.Completed)
		if result.Error != nil {
			returnInfo.ErrorExist = true
		} else {
			returnInfo.Successful = true
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}

func DeleteBenchmark(globalDB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		returnInfo := struct {
			Successful bool
			ErrorExist bool
		}{}
		benchmarkID, err := strconv.ParseUint(params["benchmarkID"], 10, 64)
		if err != nil {
			fmt.Println(err)
			returnInfo.ErrorExist = true
		}

		result := globalDB.Delete(&Benchmark{}, benchmarkID)
		if result.Error != nil {
			returnInfo.ErrorExist = true
			fmt.Printf("Error in DeleteGoal with deleting BenchmarkID:%d\n", benchmarkID)
		} else {
			returnInfo.Successful = true
		}

		json.NewEncoder(w).Encode(returnInfo)
	}
}
