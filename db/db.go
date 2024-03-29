package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	// For postgres connection
	_ "github.com/lib/pq"
)

// Conn exposes a global variable referencing the database connection
var Conn *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, dbpassword, dbname)
	Conn, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	if err = Conn.Ping(); err != nil {
		panic(err)
	}
}
