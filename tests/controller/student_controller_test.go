package controller

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/model/web"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func TestStudentControllerCreateSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webStudent.StudentCreateRequest{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7})

	result, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPost, "students")
	expected := &web.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data: map[string]interface{}{
			"id":       float64(1),
			"name":     "Hisyam",
			"email":    "hisyam@email.com",
			"active":   true,
			"semester": float64(7),
		},
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerCreateFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webStudent.StudentCreateRequest{})

	result, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPost, "students")
	expected := &web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
	}

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected.Code, result.Code)
	assert.Equal(t, expected.Status, result.Status)
}

func TestStudentControllerGetByIdSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := tests.PopulateStudentAndClassTable()
	studentRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1")
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: map[string]interface{}{
			"id":       float64(students[0].Id),
			"name":     students[0].Name,
			"email":    students[0].Email,
			"active":   students[0].Active,
			"semester": float64(students[0].Semester),
		},
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetByIdFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 1 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerUpdateSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := tests.PopulateStudentAndClassTable()
	studentRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webStudent.StudentUpdateRequest{
		Id:       students[0].Id,
		Name:     "update",
		Email:    "update@email.com",
		Active:   false,
		Semester: 5,
	})

	result, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPatch, "students/1")
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: map[string]interface{}{
			"id":       float64(1),
			"name":     "update",
			"email":    "update@email.com",
			"active":   false,
			"semester": float64(5),
		},
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerUpdateFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webStudent.StudentUpdateRequest{
		Id:       4,
		Name:     "update",
		Email:    "update@email.com",
		Active:   false,
		Semester: 5,
	})

	result, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPatch, "students/1")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 4 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerDeleteSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := tests.SetupTestRouter(db)
	tests.PopulateStudentAndClassTable()

	result, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodDelete, "students/1")
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerDeleteFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodDelete, "students/1")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 1 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetAll(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := tests.PopulateStudentAndClassTable()
	studentRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students")
	expectedData := []interface{}{}
	for _, val := range students {
		expectedData = append(expectedData, map[string]interface{}{
			"id": float64(val.Id), "name": val.Name, "email": val.Email, "active": val.Active, "semester": float64(val.Semester),
		})
	}
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   expectedData,
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetClassesById(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	_, classes := tests.PopulateEnrolledClassTable()
	studentRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1/classes")
	expectedData := []interface{}{}
	expectedData = append(expectedData, map[string]interface{}{
		"id": float64(classes[0].Id), "name": classes[0].Name, "startAt": classes[0].StartAt, "endAt": classes[0].EndAt,
	})
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   expectedData,
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetClassesByIdFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1/classes")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 1 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}
