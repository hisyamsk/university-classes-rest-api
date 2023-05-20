package repository

import (
	"context"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/repository/student"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestStudentRepositoryFindAll(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	expected := []*entity.Student{
		{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7},
		{Name: "Setiadi", Email: "setiadi@email.com", Active: true, Semester: 3},
		{Name: "Kurniawan", Email: "kurniawan@email.com", Active: false, Semester: 5},
	}
	studentRepository := student.NewStudentRepository()
	for _, val := range expected {
		studentRepository.Save(context.Background(), tx, val)
	}

	result := studentRepository.FindAll(context.Background(), tx)

	assert.Equal(t, expected, result)
}

func TestStudentRepositoryFind(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	expected := &entity.Student{Name: "John Doe", Email: "johndoe@email,com", Active: true, Semester: 6}
	studentRepository := student.NewStudentRepository()
	newStudent := studentRepository.Save(context.Background(), tx, expected)

	result, _ := studentRepository.FindById(context.Background(), tx, newStudent.Id)

	assert.Equal(t, expected, result)
}

func TestStudentRepositorySave(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	newStudent := &entity.Student{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 8}
	expected := &entity.Student{Id: 1, Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 8}
	studentRepository := student.NewStudentRepository()

	result := studentRepository.Save(context.Background(), tx, newStudent)

	assert.Equal(t, expected, result)
}

func TestStudentRepositoryUpdate(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	newStudent := &entity.Student{Id: 1, Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 8}
	studentRepository := student.NewStudentRepository()
	createdStudent := studentRepository.Save(context.Background(), tx, newStudent)
	createdStudent.Name = "Kurniawan"
	createdStudent.Active = false
	expected := &entity.Student{Id: 1, Name: "Kurniawan", Email: "hisyam@email.com", Active: false, Semester: 8}

	result := studentRepository.Update(context.Background(), tx, createdStudent)

	assert.Equal(t, expected, result)
}

func TestStudentRepositoryDelete(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	newStudent := &entity.Student{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 8}
	studentRepository := student.NewStudentRepository()
	createdStudent := studentRepository.Save(context.Background(), tx, newStudent)

	studentRepository.Delete(context.Background(), tx, createdStudent.Id)
	_, err := studentRepository.FindById(context.Background(), tx, createdStudent.Id)

	assert.NotNil(t, err)
}
