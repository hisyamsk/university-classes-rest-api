package repository

import (
	"context"
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
