package service

import (
	"context"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/model/web/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func TestEnrolledClassServiceCreate(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	tests.PopulateStudentAndClassTable()

	enrolledClassService := tests.NewTestEnrolledClassService(db)
	enrolledClassRequest := &enrolled_class.EnrolledClassRequest{StudentId: 1, ClassId: 1}

	assert.NotPanics(t, func() {
		enrolledClassService.Create(context.Background(), enrolledClassRequest)
	})
}

func TestEnrolledClassServiceDelete(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	tests.PopulateEnrolledClassTable()

	enrolledClassService := tests.NewTestEnrolledClassService(db)
	enrolledClassRequest := &enrolled_class.EnrolledClassRequest{StudentId: 1, ClassId: 1}

	assert.NotPanics(t, func() {
		enrolledClassService.Delete(context.Background(), enrolledClassRequest)
	})
}

func TestEnrolledClassServiceFind(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	tests.PopulateEnrolledClassTable()

	enrolledClassService := tests.NewTestEnrolledClassService(db)
	enrolledClassRequest := &enrolled_class.EnrolledClassRequest{StudentId: 1, ClassId: 1}

	assert.NotPanics(t, func() {
		enrolledClassService.Find(context.Background(), enrolledClassRequest)
	})
}
