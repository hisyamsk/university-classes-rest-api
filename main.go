package main

import (
	"fmt"

	"github.com/hisyamsk/university-classes-rest-api/app/db"
	_ "github.com/lib/pq"
)

func main() {
	db := db.NewDBConnection()

	fmt.Println(db)
}
