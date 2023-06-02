package tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/db"
	classController "github.com/hisyamsk/university-classes-rest-api/controller/class"
	enrolledClassController "github.com/hisyamsk/university-classes-rest-api/controller/enrolled_class"
	studentController "github.com/hisyamsk/university-classes-rest-api/controller/student"
	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/middleware"
	"github.com/hisyamsk/university-classes-rest-api/model/web"
	"github.com/hisyamsk/university-classes-rest-api/repository/class"
	classRepository "github.com/hisyamsk/university-classes-rest-api/repository/class"
	"github.com/hisyamsk/university-classes-rest-api/repository/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/repository/student"
	studentRepository "github.com/hisyamsk/university-classes-rest-api/repository/student"
	classService "github.com/hisyamsk/university-classes-rest-api/service/class"
	enrolledClassService "github.com/hisyamsk/university-classes-rest-api/service/enrolled_class"
	studentService "github.com/hisyamsk/university-classes-rest-api/service/student"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const API_URL = "http://localhost:8000/api"

func SetupTestDB() (*sql.Tx, *sql.DB) {
	helper.LoadDotenv("../../.env")
	database := db.NewDBConnection(app.DbNameTest)
	tx, err := database.Begin()
	helper.PanicIfError(err)

	return tx, database
}

func CleanUpTest(tx *sql.Tx, db *sql.DB) {
	helper.CommitOrRollback(tx)
	_, err := db.Exec("TRUNCATE enrolled_class, student, class RESTART IDENTITY")
	helper.PanicIfError(err)
}

func PopulateStudentAndClassTable() ([]*entity.Student, []*entity.Class) {
	tx, db := SetupTestDB()
	_, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students := []*entity.Student{
		{Name: "Hisyam", Email: "hisyam@email.com", Active: true, Semester: 7},
		{Name: "Setiadi", Email: "setiadi@email.com", Active: false, Semester: 5},
		{Name: "Kurniawan", Email: "kurniawan@email.com", Active: true, Semester: 5},
	}
	classes := []*entity.Class{
		{Name: "Algorithm and Data Structures", StartAt: "07:00:00", EndAt: "09:00:00"},
		{Name: "Linear Algebra", StartAt: "07:00:00", EndAt: "09:00:00"},
		{Name: "Discrete Math", StartAt: "07:00:00", EndAt: "09:00:00"},
	}
	studentRepository := student.NewStudentRepository()
	classRepository := class.NewClassRepositoryImpl()

	for _, val := range students {
		studentRepository.Save(context.Background(), tx, val)
	}
	for _, val := range classes {
		classRepository.Save(context.Background(), tx, val)
	}

	return students, classes
}

func PopulateEnrolledClassTable() ([]*entity.Student, []*entity.Class) {
	tx, db := SetupTestDB()
	_, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students, classes := PopulateStudentAndClassTable()
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()
	for i := range students {
		enrolledClassRepository.Save(context.Background(), tx, &entity.EnrolledClass{StudentId: students[i].Id, ClassId: classes[i].Id})
	}

	return students, classes
}

func NewTestClassService(db *sql.DB) classService.ClassService {
	validate := validator.New()
	classRepository := classRepository.NewClassRepositoryImpl()
	classService := classService.NewClassService(db, classRepository, validate)

	return classService
}

func NewTestStudentService(db *sql.DB) studentService.StudentService {
	validate := validator.New()
	studentRepository := studentRepository.NewStudentRepository()
	studentService := studentService.NewStudentService(studentRepository, db, validate)

	return studentService
}

func NewTestEnrolledClassService(db *sql.DB) enrolledClassService.EnrolledClassService {
	validate := validator.New()
	enrolledClassRepository := enrolled_class.NewEnrolledClassRepository()
	enrolledClassService := enrolledClassService.NewEnrolledClassService(db, enrolledClassRepository, validate)

	return enrolledClassService
}

func NewTestStudentController(db *sql.DB) studentController.StudentController {
	studentService := NewTestStudentService(db)
	studentController := studentController.NewStudentController(studentService)

	return studentController
}

func NewTestClassController(db *sql.DB) classController.ClassController {
	classService := NewTestClassService(db)
	classController := classController.NewClassController(classService)

	return classController
}

func NewTestEnrolledClassController(db *sql.DB) enrolledClassController.EnrolledClassController {
	enrolledClassService := NewTestEnrolledClassService(db)
	enrolledClassController := enrolledClassController.NewEnrolledClassController(enrolledClassService)

	return enrolledClassController
}

func NewTestEnrolledClassMiddleware(db *sql.DB) middleware.EnrolledClassMiddleware {
	classService := NewTestClassService(db)
	studentService := NewTestStudentService(db)
	enrolledClassMiddleware := middleware.NewEnrolledClassMiddleware(classService, studentService)

	return enrolledClassMiddleware
}

func SetupTestRouter(db *sql.DB) *httprouter.Router {
	studentController := NewTestStudentController(db)
	classController := NewTestClassController(db)
	enrolledClassController := NewTestEnrolledClassController(db)
	enrolledClassMiddleware := NewTestEnrolledClassMiddleware(db)
	routerHandler := &app.RouterHandler{
		StudentController:       studentController,
		ClassController:         classController,
		EnrolledClassController: enrolledClassController,
		EnrolledClassMiddleware: enrolledClassMiddleware,
	}
	router := app.NewRouter(routerHandler)

	return router
}

func SetupRequestAndRecorder(router *httprouter.Router, body []byte, method, endpoint string) (*web.WebResponse, *http.Response) {
	requestBody := strings.NewReader(string(body))
	request := httptest.NewRequest(method, fmt.Sprintf("%s/%s", API_URL, endpoint), requestBody)
	request.Header.Add("X-API-Key", os.Getenv("API_KEY_SECRET"))
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	authMiddleware := middleware.NewAuthMiddleware(router)

	authMiddleware.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)
	var result *web.WebResponse
	err := json.Unmarshal(responseBody, &result)
	helper.PanicIfError(err)

	return result, response
}
