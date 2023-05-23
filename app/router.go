package app

import (
	studentController "github.com/hisyamsk/university-classes-rest-api/controller/student"
	"github.com/hisyamsk/university-classes-rest-api/exception"
	"github.com/julienschmidt/httprouter"
)

type RouterHandler struct {
	StudentController studentController.StudentController
}

func NewRouter(handler *RouterHandler) *httprouter.Router {
	router := httprouter.New()

	// /students
	router.GET("/api/students", handler.StudentController.GetAll)
	router.POST("/api/students", handler.StudentController.Create)
	router.GET("/api/students/:id", handler.StudentController.GetById)
	router.PATCH("/api/students/:id", handler.StudentController.Update)
	router.DELETE("/api/students/:id", handler.StudentController.Delete)
	router.GET("/api/students/:id/classes", handler.StudentController.GetClassesById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
