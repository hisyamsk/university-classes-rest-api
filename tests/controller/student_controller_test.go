package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/web"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func setupTestStudentRouter(db *sql.DB) http.Handler {
	studentController := tests.NewTestStudentController(db)
	routerHandler := &app.RouterHandler{
		StudentController: studentController,
	}
	router := app.NewRouter(routerHandler)

	return router
}

func TestStudentControllerCreateSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := setupTestStudentRouter(db)
	body, _ := json.Marshal(&webStudent.StudentCreateRequest{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7})

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPost, "students")
	var result *web.WebResponse
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
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerCreateFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := setupTestStudentRouter(db)
	body, _ := json.Marshal(&webStudent.StudentCreateRequest{})

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPost, "students")
	var result *web.WebResponse
	expected := &web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected.Code, result.Code)
	assert.Equal(t, expected.Status, result.Status)
}

func TestStudentControllerGetByIdSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := tests.PopulateStudentAndClassTable()
	studentRouter := setupTestStudentRouter(db)

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1")
	var result *web.WebResponse
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
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetByIdFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := setupTestStudentRouter(db)

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1")
	var result *web.WebResponse
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 1 was not found",
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerUpdateSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := tests.PopulateStudentAndClassTable()
	studentRouter := setupTestStudentRouter(db)
	body, _ := json.Marshal(&webStudent.StudentUpdateRequest{
		Id:       students[0].Id,
		Name:     "update",
		Email:    "update@email.com",
		Active:   false,
		Semester: 5,
	})

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPatch, "students/1")
	var result *web.WebResponse
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
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerUpdateFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := setupTestStudentRouter(db)
	body, _ := json.Marshal(&webStudent.StudentUpdateRequest{
		Id:       4,
		Name:     "update",
		Email:    "update@email.com",
		Active:   false,
		Semester: 5,
	})

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, body, http.MethodPatch, "students/1")
	var result *web.WebResponse
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 4 was not found",
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerDeleteSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := setupTestStudentRouter(db)
	tests.PopulateStudentAndClassTable()

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodDelete, "students/1")
	var result *web.WebResponse
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerDeleteFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := setupTestStudentRouter(db)

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodDelete, "students/1")
	var result *web.WebResponse
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 1 was not found",
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetAll(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	students, _ := tests.PopulateStudentAndClassTable()
	studentRouter := setupTestStudentRouter(db)

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students")
	var result *web.WebResponse
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
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetClassesById(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	_, classes := tests.PopulateEnrolledClassTable()
	studentRouter := setupTestStudentRouter(db)

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1/classes")
	var result *web.WebResponse
	expectedData := []interface{}{}
	expectedData = append(expectedData, map[string]interface{}{
		"id": float64(classes[0].Id), "name": classes[0].Name, "startAt": classes[0].StartAt, "endAt": classes[0].EndAt,
	})
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   expectedData,
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestStudentControllerGetClassesByIdFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	studentRouter := setupTestStudentRouter(db)

	responseBody, response := tests.SetupRequestAndRecorder(studentRouter, nil, http.MethodGet, "students/1/classes")
	var result *web.WebResponse
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Student with id: 1 was not found",
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}
