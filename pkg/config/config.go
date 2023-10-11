package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func Connect() {
	//Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load env vars: %s", err)
		os.Exit(1)
	}

	//Get db credentials from environment variable
	dbToken := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	
	var dbUrl = "libsql://"+dbName+".turso.io?authToken="+dbToken

	d, err := sql.Open("libsql", dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}

	db = d
}

func GetDB() *sql.DB {
	return db
}