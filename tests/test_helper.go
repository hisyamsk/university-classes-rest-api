package tests

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/db"
	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/repository/class"
	"github.com/hisyamsk/university-classes-rest-api/repository/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/repository/student"
	_ "github.com/lib/pq"
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

func PopulateStudentAndClassTable() ([]*entity.Student, []*entity.Class) {
	tx, db := SetupTestDB()
	_, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students := []*entity.Student{
		{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7},
		{Name: "Setiadi", Email: "setiadi@email.com", Active: false, Semester: 5},
		{Name: "Kurniawan", Email: "kurniawan@email.com", Active: true, Semester: 5},
	}
	classes := []*entity.Class{
		{Name: "Algorithm and Data Structures", StartAt: "07:00:00", EndAt: "09:00:00"},
		{Name: "Linear Algebra", StartAt: "07:00:00", EndAt: "09:00:00"},
		{Name: "Discrete Math", StartAt: "07:00:00", EndAt: "09:00:00"},
	}
	studentRepository := student.NewStudentRepository()
	classRepository := class.NewClassRepositoryImpl()

	for _, val := range students {
		studentRepository.Save(context.Background(), tx, val)
	}
	for _, val := range classes {
		classRepository.Save(context.Background(), tx, val)
	}

	return students, classes
}

func PopulateEnrolledClassTable() ([]*entity.Student, []*entity.Class) {
	tx, db := SetupTestDB()
	_, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students, classes := PopulateStudentAndClassTable()
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()
	for i := range students {
		enrolledClassRepository.Save(context.Background(), tx, &entity.EnrolledClass{StudentId: students[i].Id, ClassId: classes[i].Id})
	}

	return students, classes
}
