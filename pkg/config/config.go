package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type DbConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

func Connect() {

	dbConfig := getDbConfig()

	//construct url with config
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	d, err := gorm.Open("mysql", dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s\n", dbUrl, err)
		os.Exit(1)
	}

	db = d

}

func getDbConfig() DbConfig {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load env vars: %s\n", err)
		os.Exit(1)
	}

	// Get db credentials from environment variable
	dbusername := os.Getenv("dbusername")
	dbuserpassword := os.Getenv("dbuserpassword")
	dbname := os.Getenv("dbname")

	// ensures values that are required are not empty strings!
	if dbusername == "" || dbuserpassword == "" || dbname == "" {
		fmt.Fprintf(os.Stderr, "One or more required environment variables (dbusername, dbuserpassword, dbname) are not set.\n")
		os.Exit(1)
	}

	// Get DB_HOST from environment variable with default value
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // Set default host
	}

	// Get and validate DB_PORT
	dbPortEnv := os.Getenv("DB_PORT")
	dbPort := 9000 // Default port
	if dbPortEnv != "" {
		var err error
		dbPort, err = strconv.Atoi(dbPortEnv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid port: %s, using default port: %d\n", dbPortEnv, dbPort)
			dbPort = 9000 // Reset to default value but don't exit.
		}
	}

	return DbConfig{
		User:     dbusername,
		Password: dbuserpassword,
		DBName:   dbname,
		Host:     dbHost,
		Port:     dbPort,
	}
}

func GetDB() *gorm.DB {
	return db
}
