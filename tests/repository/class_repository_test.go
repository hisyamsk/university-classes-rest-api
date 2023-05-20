package repository

import (
	"context"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/repository/class"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func TestClassRepositoryFind(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	expected := &entity.Class{Name: "Discrete Math", StartAt: "07:00:00", EndAt: "09:00:00"}
	classRepository := class.NewClassRepositoryImpl()
	createdClass := classRepository.Save(context.Background(), tx, expected)

	result, _ := classRepository.FindById(context.Background(), tx, createdClass.Id)

	assert.Equal(t, expected, result)
}

func TestClassRepositoryFindAll(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	expected := []*entity.Class{
		{Name: "Algorithm and Data Structures", StartAt: "08:00:00", EndAt: "10:00:00"},
		{Name: "Discrete Math", StartAt: "09:00:00", EndAt: "11:00:00"},
		{Name: "Linear Algebra", StartAt: "13:00:00", EndAt: "15:00:00"},
	}
	classRepository := class.NewClassRepositoryImpl()
	for _, val := range expected {
		classRepository.Save(context.Background(), tx, val)
	}

	result := classRepository.FindAll(context.Background(), tx)

	assert.Equal(t, expected, result)
}

func TestClassRepositorySave(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	newClass := &entity.Class{Name: "Linear Algebra", StartAt: "13:00:00", EndAt: "15:00:00"}
	expected := &entity.Class{Id: 1, Name: "Linear Algebra", StartAt: "13:00:00", EndAt: "15:00:00"}
	classRepository := class.NewClassRepositoryImpl()

	result := classRepository.Save(context.Background(), tx, newClass)

	assert.Equal(t, expected, result)
}

func TestClassRepositoryUpdate(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	newClass := &entity.Class{Name: "Linear Algebra", StartAt: "13:00:00", EndAt: "15:00:00"}
	expected := &entity.Class{Id: 1, Name: "Discrete Math", StartAt: "15:00:00", EndAt: "17:00:00"}
	classRepository := class.NewClassRepositoryImpl()
	createdClass := classRepository.Save(context.Background(), tx, newClass)
	createdClass.Name = "Discrete Math"
	createdClass.StartAt = "15:00:00"
	createdClass.EndAt = "17:00:00"

	result := classRepository.Update(context.Background(), tx, createdClass)

	assert.Equal(t, expected, result)
}

func TestClassRepositoryDelete(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	newClass := &entity.Class{Name: "Linear Algebra", StartAt: "13:00:00", EndAt: "15:00:00"}
	classRepository := class.NewClassRepositoryImpl()
	createdClass := classRepository.Save(context.Background(), tx, newClass)

	classRepository.Delete(context.Background(), tx, createdClass.Id)
	_, err := classRepository.FindById(context.Background(), tx, createdClass.Id)

	assert.NotNil(t, err)
}
