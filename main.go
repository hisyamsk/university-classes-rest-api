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

	handler := server.InitializeHandler(dbName)
	serverHandler := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	fmt.Println("listening on port", addr)
	err := serverHandler.ListenAndServe()
	helper.PanicIfError(err)
}
