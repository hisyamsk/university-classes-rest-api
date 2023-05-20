package tests

import (
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/db"
	"github.com/hisyamsk/university-classes-rest-api/helper"
)

func SetupTestDB() (*sql.Tx, *sql.DB) {
	database := db.NewDBConnection(app.DbNameTest)
	tx, err := database.Begin()
	helper.PanicIfError(err)

	return tx, database
}

func CleanUpTest(tx *sql.Tx, db *sql.DB) {
	helper.CommitOrRollback(tx)
	_, err := db.Exec("TRUNCATE enrolled_class, student, class RESTART IDENTITY")
	helper.PanicIfError(err)
}
