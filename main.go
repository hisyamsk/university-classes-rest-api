package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/server"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	dbName := app.DbName
	addr := os.Getenv("APP_ADDRESS")

	router := server.InitializeServer(dbName)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Printf("listening on %s", addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
