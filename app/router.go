package app

import (
	classController "github.com/hisyamsk/university-classes-rest-api/controller/class"
	studentController "github.com/hisyamsk/university-classes-rest-api/controller/student"
	"github.com/hisyamsk/university-classes-rest-api/exception"
	"github.com/julienschmidt/httprouter"
)

type RouterHandler struct {
	StudentController studentController.StudentController
	ClassController   classController.ClassController
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

	// /class
	router.GET("/api/classes", handler.ClassController.GetAll)
	router.POST("/api/classes", handler.ClassController.Create)
	router.GET("/api/classes/:id", handler.ClassController.GetById)
	router.PATCH("/api/classes/:id", handler.ClassController.Update)
	router.DELETE("/api/classes/:id", handler.ClassController.Delete)
	router.GET("/api/classes/:id/students", handler.ClassController.GetStudentsById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
