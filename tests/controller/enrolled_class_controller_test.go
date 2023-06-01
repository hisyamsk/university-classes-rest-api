package controller

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/model/web"
	"github.com/hisyamsk/university-classes-rest-api/model/web/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func TestEnrolledClassControllerCreate(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, classes := tests.PopulateStudentAndClassTable()
	router := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&enrolled_class.EnrolledClassRequest{StudentId: students[0].Id, ClassId: classes[0].Id})

	result, response := tests.SetupRequestAndRecorder(router, body, http.MethodPost, "enrolled-class")
	expected := &web.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestEnrolledClassControllerDeleteSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, classes := tests.PopulateEnrolledClassTable()
	router := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&enrolled_class.EnrolledClassRequest{StudentId: students[0].Id, ClassId: classes[0].Id})

	result, response := tests.SetupRequestAndRecorder(router, body, http.MethodDelete, "enrolled-class")
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestEnrolledClassControllerDeleteFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	router := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&enrolled_class.EnrolledClassRequest{StudentId: 99, ClassId: 99})

	result, response := tests.SetupRequestAndRecorder(router, body, http.MethodDelete, "enrolled-class")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "studentId or classId was not found!",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}
