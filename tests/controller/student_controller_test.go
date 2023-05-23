package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
	requestBody := strings.NewReader(string(body))
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s/students", tests.API_URL), requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	studentRouter.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
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
	requestBody := strings.NewReader(string(body))
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s/students", tests.API_URL), requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	studentRouter.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
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
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/students/1", tests.API_URL), nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	studentRouter.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
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
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/students/1", tests.API_URL), nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	studentRouter.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
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
