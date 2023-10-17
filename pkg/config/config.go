package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	//Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load env vars: %s", err)
		os.Exit(1)
	}

	// Get db credentials from environment variable
	dbname := os.Getenv("dbname")
	dbpassword := os.Getenv("dbpassword")
	
	var dbUrl = dbname+":"+dbpassword+"@tcp(localhost:9000)/moviesdb?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}