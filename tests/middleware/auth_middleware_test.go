package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/middleware"
	"github.com/hisyamsk/university-classes-rest-api/model/web"
	"github.com/hisyamsk/university-classes-rest-api/tests"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddlewareUnauthorized(t *testing.T) {
	tx, db := tests.SetupTestDB()
	defer tests.CleanUpTest(tx, db)
	router := tests.SetupTestRouter(db)
	authMiddleware := middleware.NewAuthMiddleware(router)

	request := httptest.NewRequest(http.MethodGet, tests.API_URL+"/students", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "WRONG SECRET")
	recorder := httptest.NewRecorder()

	authMiddleware.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
	var result *web.WebResponse
	expected := &web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: http.StatusText(http.StatusUnauthorized),
	}
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.Equal(t, expected, result)
}
