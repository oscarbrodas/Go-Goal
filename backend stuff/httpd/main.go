package main

// Parts Taken From Dobra's Code

import (
	"log"
	"net/http"

	"go-goal/httpd/handler"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

// mySQL DSN string is no longer needed since we are using sqlite
// const DSN string = "root:password@tcp(127.0.0.1:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("database failed to open")
	}
	globalDB = db

	// MAKE SURE TO AUTOMIGRATE BEFORE INITALIZEROUTER
	globalDB.AutoMigrate(&handler.User{})
	globalDB.AutoMigrate(&handler.Goal{})
	globalDB.AutoMigrate(&handler.Friend{})

	host := "127.0.0.1:9000" // :9000 is the port
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}
