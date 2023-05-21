package repository

import (
	"context"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/repository/class"
	"github.com/hisyamsk/university-classes-rest-api/repository/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/repository/student"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func populateStudentAndClassTable() ([]*entity.Student, []*entity.Class) {
	tx, db := tests.SetupTestDB()
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

func TestEnrolledClassRepositorySave(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	populateStudentAndClassTable()
	newEnrolledClass := &entity.EnrolledClass{StudentId: 1, ClassId: 1}
	expected := &entity.EnrolledClass{Id: 1, StudentId: 1, ClassId: 1}
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()

	result := enrolledClassRepository.Save(context.Background(), tx, newEnrolledClass)

	assert.Equal(t, expected, result)
}

func TestEnrolledClassRepositoryFindByStudentId(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	_, classes := populateStudentAndClassTable()
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()
	newEnrolledClass := &entity.EnrolledClass{StudentId: 1, ClassId: 1}
	enrolledClassRepository.Save(context.Background(), tx, newEnrolledClass)
	expected := []*entity.Class{classes[0]}

	result := enrolledClassRepository.FindByStudentId(context.Background(), tx, newEnrolledClass.StudentId)

	assert.Equal(t, expected, result)
}

func TestEnrolledClassRepositoryFindByClassId(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := populateStudentAndClassTable()
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()
	newEnrolledClass := &entity.EnrolledClass{StudentId: 1, ClassId: 1}
	enrolledClassRepository.Save(context.Background(), tx, newEnrolledClass)
	expected := []*entity.Student{students[0]}

	result := enrolledClassRepository.FindByClassId(context.Background(), tx, newEnrolledClass.ClassId)

	assert.Equal(t, expected, result)
}
