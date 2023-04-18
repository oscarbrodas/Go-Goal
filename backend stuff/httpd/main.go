package main

// Parts Taken From Dobra's Code

import (
	"log"
	"net/http"

	"go-goal/httpd/handler"

	"github.com/joho/godotenv"

	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var globalDB *gorm.DB
var globalUploader *manager.Uploader
var globalDownloader *manager.Downloader

// mySQL DSN string is no longer needed since we are using sqlite
// const DSN string = "root:password@tcp(127.0.0.1:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup SQlite DB
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("database failed to open")
	}
	globalDB = db

	// MAKE SURE TO AUTOMIGRATE BEFORE INITALIZEROUTER
	globalDB.AutoMigrate(&handler.User{})
	globalDB.AutoMigrate(&handler.Goal{})
	globalDB.AutoMigrate(&handler.Friend{})
	globalDB.AutoMigrate(&handler.Benchmark{})

	// Setup s3 uploader
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	globalUploader = manager.NewUploader(client)
	globalDownloader = manager.NewDownloader(client)

	// Start Router
	host := "127.0.0.1:9000" // :9000 is the port
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}

}
