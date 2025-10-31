package db

import (
	"database/sql"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB
	
func Init() {
	_ = godotenv.Load()
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN environment variable not set")
	}
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
}
