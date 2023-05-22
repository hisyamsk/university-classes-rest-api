package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	web "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	studentRepository "github.com/hisyamsk/university-classes-rest-api/repository/student"
	"github.com/hisyamsk/university-classes-rest-api/service/student"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func newStudentService(db *sql.DB) student.StudentService {
	validate := validator.New()
	studentRepository := studentRepository.NewStudentRepository()
	studentService := student.NewStudentService(studentRepository, db, validate)

	return studentService
}

func TestStudentServiceFindById(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	studentService := newStudentService(db)
	studentRequest := &web.StudentCreateRequest{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7}
	createdStudent := studentService.Create(context.Background(), studentRequest)
	expected := &web.StudentResponse{Id: 1, Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7}

	result, err := studentService.FindById(context.Background(), createdStudent.Id)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestStudentServiceFindAll(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	studentService := newStudentService(db)
	studentRequest := &web.StudentCreateRequest{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7}
	createdStudent := studentService.Create(context.Background(), studentRequest)
	expected := []*web.StudentResponse{}
	expected = append(expected, createdStudent)

	result := studentService.FindAll(context.Background())

	assert.Equal(t, expected, result)
}

func TestStudentServiceSave(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	studentService := newStudentService(db)
	studentRequest := &web.StudentCreateRequest{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7}
	expected := &web.StudentResponse{Id: 1, Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7}

	result := studentService.Create(context.Background(), studentRequest)

	assert.Equal(t, expected, result)
}

func TestStudentServiceUpdate(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	studentService := newStudentService(db)
	studentRequest := &web.StudentCreateRequest{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7}
	studentService.Create(context.Background(), studentRequest)
	updateStudent := &web.StudentUpdateRequest{Id: 1, Name: "Kurniawan", Email: "kurniawan@email.com", Active: true, Semester: 7}
	expected := &web.StudentResponse{Id: 1, Name: "Kurniawan", Email: "kurniawan@email.com", Active: true, Semester: 7}

	result := studentService.Update(context.Background(), updateStudent)

	assert.Equal(t, expected, result)
}

func TestStudentServiceDelete(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	studentService := newStudentService(db)
	studentRequest := &web.StudentCreateRequest{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7}
	createdStudent := studentService.Create(context.Background(), studentRequest)

	studentService.Delete(context.Background(), createdStudent.Id)
	result, err := studentService.FindById(context.Background(), createdStudent.Id)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestStudentServiceFindClasses(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	studentService := newStudentService(db)
	_, classes := tests.PopulateEnrolledClassTable()
	expected := helper.ToClassesResponse([]*entity.Class{classes[0]})

	result := studentService.FindClasses(context.Background(), 1)

	assert.Equal(t, expected, result)
}
