package main

import (
	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/db"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	dbName := app.DbName

	db.NewDBConnection(dbName)
}
