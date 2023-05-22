package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	webClass "github.com/hisyamsk/university-classes-rest-api/model/web/class"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	classRepository "github.com/hisyamsk/university-classes-rest-api/repository/class"
	classService "github.com/hisyamsk/university-classes-rest-api/service/class"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func newClassService(db *sql.DB) classService.ClassService {
	validate := validator.New()
	classRepository := classRepository.NewClassRepositoryImpl()
	classService := classService.NewClassService(db, classRepository, validate)

	return classService
}

func TestClassServiceCreate(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	classService := newClassService(db)
	classRequest := &webClass.ClassCreateRequest{Name: "Graph Theory", StartAt: "07:00:00", EndAt: "09:00:00"}
	expected := &webClass.ClassResponse{Id: 1, Name: "Graph Theory", StartAt: "07:00:00", EndAt: "09:00:00"}

	result := classService.Create(context.Background(), classRequest)

	assert.Equal(t, expected, result)
}

func TestClassServiceFindById(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	classService := newClassService(db)
	classRequest := &webClass.ClassCreateRequest{Name: "Graph Theory", StartAt: "07:00:00", EndAt: "09:00:00"}
	createdClass := classService.Create(context.Background(), classRequest)
	expected := &webClass.ClassResponse{Id: 1, Name: "Graph Theory", StartAt: "07:00:00", EndAt: "09:00:00"}

	result, err := classService.FindById(context.Background(), createdClass.Id)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestClassServiceUpdate(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	classService := newClassService(db)
	classCreateRequest := &webClass.ClassCreateRequest{Name: "Graph Theory", StartAt: "07:00:00", EndAt: "09:00:00"}
	createdClass := classService.Create(context.Background(), classCreateRequest)
	classUpdateRequest := &webClass.ClassUpdateRequest{Id: createdClass.Id, Name: createdClass.Name, StartAt: "09:00:00", EndAt: "11:00:00"}
	expected := &webClass.ClassResponse{Id: 1, Name: "Graph Theory", StartAt: "09:00:00", EndAt: "11:00:00"}

	result := classService.Update(context.Background(), classUpdateRequest)

	assert.Equal(t, expected, result)
}

func TestClassServiceDelete(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	classService := newClassService(db)
	classCreateRequest := &webClass.ClassCreateRequest{Name: "Graph Theory", StartAt: "07:00:00", EndAt: "09:00:00"}
	createdClass := classService.Create(context.Background(), classCreateRequest)

	classService.Delete(context.Background(), createdClass.Id)
	class, err := classService.FindById(context.Background(), createdClass.Id)

	assert.Nil(t, class)
	assert.NotNil(t, err)
}

func TestClassServiceFindAll(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)

	classService := newClassService(db)
	classCreateRequest := &webClass.ClassCreateRequest{Name: "Graph Theory", StartAt: "07:00:00", EndAt: "09:00:00"}
	createdClass := classService.Create(context.Background(), classCreateRequest)
	expected := []*webClass.ClassResponse{createdClass}

	result := classService.FindAll(context.Background())

	assert.Equal(t, expected, result)
}

func TestClassServiceFindStundentsById(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := tests.PopulateEnrolledClassTable()

	classService := newClassService(db)
	expected := []*webStudent.StudentResponse{helper.ToStudentResponse(students[0])}
	result := classService.FindStudentsById(context.Background(), 1)

	assert.Equal(t, expected, result)
}
