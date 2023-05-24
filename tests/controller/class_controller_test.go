package controller

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/model/web"
	webClass "github.com/hisyamsk/university-classes-rest-api/model/web/class"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func TestClassControllerCreateSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webClass.ClassCreateRequest{Name: "Statistics", StartAt: "09:30:00", EndAt: "10:50:00"})

	result, response := tests.SetupRequestAndRecorder(classRouter, body, http.MethodPost, "classes")
	expected := &web.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data: map[string]interface{}{
			"id":      float64(1),
			"name":    "Statistics",
			"startAt": "09:30:00",
			"endAt":   "10:50:00",
		},
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerCreateFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webClass.ClassCreateRequest{})

	result, response := tests.SetupRequestAndRecorder(classRouter, body, http.MethodPost, "classes")
	expected := &web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
	}

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected.Code, result.Code)
	assert.Equal(t, expected.Status, result.Status)
}

func TestClassControllerGetByIdSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	_, classes := tests.PopulateStudentAndClassTable()
	classRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(classRouter, nil, http.MethodGet, "classes/1")
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: map[string]interface{}{
			"id":      float64(classes[0].Id),
			"name":    classes[0].Name,
			"startAt": classes[0].StartAt,
			"endAt":   classes[0].EndAt,
		},
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerGetByIdFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(classRouter, nil, http.MethodGet, "classes/1")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Class with id: 1 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerUpdateSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	_, classes := tests.PopulateStudentAndClassTable()
	classRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webClass.ClassUpdateRequest{Id: classes[0].Id, Name: "Statistics", StartAt: classes[0].StartAt, EndAt: classes[0].EndAt})

	result, response := tests.SetupRequestAndRecorder(classRouter, body, http.MethodPatch, "classes/1")
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: map[string]interface{}{
			"id":      float64(classes[0].Id),
			"name":    "Statistics",
			"startAt": classes[0].StartAt,
			"endAt":   classes[0].EndAt,
		},
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerUpdateFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)
	body, _ := json.Marshal(&webClass.ClassUpdateRequest{Id: 1, Name: "Discrete Math", StartAt: "09:00:00", EndAt: "10:30:00"})

	result, response := tests.SetupRequestAndRecorder(classRouter, body, http.MethodPatch, "classes/1")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Class with id: 1 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerDeleteSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)
	tests.PopulateStudentAndClassTable()

	result, response := tests.SetupRequestAndRecorder(classRouter, nil, http.MethodDelete, "classes/1")
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerDeleteFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(classRouter, nil, http.MethodDelete, "classes/1")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Class with id: 1 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerGetAll(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)
	_, classes := tests.PopulateEnrolledClassTable()

	result, response := tests.SetupRequestAndRecorder(classRouter, nil, http.MethodGet, "classes")
	expectedData := []any{}
	for _, class := range classes {
		expectedData = append(expectedData, map[string]any{
			"id":      float64(class.Id),
			"name":    class.Name,
			"startAt": class.StartAt,
			"endAt":   class.EndAt,
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

func TestClassControllerGetStudentsByIdSuccess(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)
	students, _ := tests.PopulateEnrolledClassTable()

	result, response := tests.SetupRequestAndRecorder(classRouter, nil, http.MethodGet, "classes/1/students")
	expectedData := []any{
		map[string]any{
			"id":       float64(students[0].Id),
			"name":     students[0].Name,
			"email":    students[0].Email,
			"active":   students[0].Active,
			"semester": float64(students[0].Semester),
		},
	}
	expected := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   expectedData,
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestClassControllerGetStudentsByIdFailed(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	classRouter := tests.SetupTestRouter(db)

	result, response := tests.SetupRequestAndRecorder(classRouter, nil, http.MethodGet, "classes/1/students")
	expected := &web.WebResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   "Class with id: 1 was not found",
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}
