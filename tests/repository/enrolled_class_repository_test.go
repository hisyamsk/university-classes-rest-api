package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/repository/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func TestEnrolledClassRepositorySave(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	tests.PopulateStudentAndClassTable()
	newEnrolledClass := &entity.EnrolledClass{StudentId: 1, ClassId: 1}
	expected := &entity.EnrolledClass{Id: 1, StudentId: 1, ClassId: 1}
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()

	result := enrolledClassRepository.Save(context.Background(), tx, newEnrolledClass)

	assert.Equal(t, expected, result)
}

func TestEnrolledClassRepositoryFind(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, classes := tests.PopulateEnrolledClassTable()
	expected := &entity.EnrolledClass{Id: 1, StudentId: students[0].Id, ClassId: classes[0].Id}
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()

	result, err := enrolledClassRepository.Find(context.Background(), tx, expected)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestEnrolledClassRepositoryDelete(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, classes := tests.PopulateEnrolledClassTable()
	enrolledClass := &entity.EnrolledClass{StudentId: students[0].Id, ClassId: classes[0].Id}
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()

	enrolledClassRepository.Delete(context.Background(), tx, enrolledClass)
	result, err := enrolledClassRepository.Find(context.Background(), tx, enrolledClass)
	expected := fmt.Errorf("studentId or classId was not found!")

	assert.Nil(t, result)
	assert.Equal(t, err, expected)
}
