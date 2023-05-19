package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/joho/godotenv"
)

func NewDBConnection() *sql.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=university_classes_db host=%s port=%s sslmode=disable", username, password, host, port)
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
