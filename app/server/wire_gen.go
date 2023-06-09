// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/db"
	class3 "github.com/hisyamsk/university-classes-rest-api/controller/class"
	enrolled_class3 "github.com/hisyamsk/university-classes-rest-api/controller/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/controller/student"
	"github.com/hisyamsk/university-classes-rest-api/middleware"
	"github.com/hisyamsk/university-classes-rest-api/repository/class"
	"github.com/hisyamsk/university-classes-rest-api/repository/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/repository/student"
	class2 "github.com/hisyamsk/university-classes-rest-api/service/class"
	enrolled_class2 "github.com/hisyamsk/university-classes-rest-api/service/enrolled_class"
	student2 "github.com/hisyamsk/university-classes-rest-api/service/student"
	"net/http"
)

// Injectors from injector.go:

func InitializeHandler(dbName string) http.Handler {
	studentRepositoryImpl := student.NewStudentRepository()
	sqlDB := db.NewDBConnection(dbName)
	validate := validator.New()
	studentServiceImpl := student2.NewStudentService(studentRepositoryImpl, sqlDB, validate)
	studentControllerImpl := controller.NewStudentController(studentServiceImpl)
	classRepositoryImpl := class.NewClassRepositoryImpl()
	classServiceImpl := class2.NewClassService(sqlDB, classRepositoryImpl, validate)
	classControllerImpl := class3.NewClassController(classServiceImpl)
	enrolledClassRepositoryImpl := enrolled_class.NewEnrolledClassRepository()
	enrolledClassServiceImpl := enrolled_class2.NewEnrolledClassService(sqlDB, enrolledClassRepositoryImpl, validate)
	enrolledClassControllerImpl := enrolled_class3.NewEnrolledClassController(enrolledClassServiceImpl)
	enrolledClassMiddlewareImpl := middleware.NewEnrolledClassMiddleware(classServiceImpl, studentServiceImpl)
	routerHandler := &app.RouterHandler{
		StudentController:       studentControllerImpl,
		ClassController:         classControllerImpl,
		EnrolledClassController: enrolledClassControllerImpl,
		EnrolledClassMiddleware: enrolledClassMiddlewareImpl,
	}
	router := app.NewRouter(routerHandler)
	handler := middleware.NewAuthMiddleware(router)
	return handler
}

// injector.go:

var classSet = wire.NewSet(class.NewClassRepositoryImpl, wire.Bind(new(class.ClassRepository), new(*class.ClassRepositoryImpl)), class2.NewClassService, wire.Bind(new(class2.ClassService), new(*class2.ClassServiceImpl)), class3.NewClassController)

var studentSet = wire.NewSet(student.NewStudentRepository, wire.Bind(new(student.StudentRepository), new(*student.StudentRepositoryImpl)), student2.NewStudentService, wire.Bind(new(student2.StudentService), new(*student2.StudentServiceImpl)), controller.NewStudentController)

var enrolledClassSet = wire.NewSet(enrolled_class.NewEnrolledClassRepository, wire.Bind(new(enrolled_class.EnrolledClassRepository), new(*enrolled_class.EnrolledClassRepositoryImpl)), enrolled_class2.NewEnrolledClassService, wire.Bind(new(enrolled_class2.EnrolledClassService), new(*enrolled_class2.EnrolledClassServiceImpl)), enrolled_class3.NewEnrolledClassController, middleware.NewEnrolledClassMiddleware, wire.Bind(new(middleware.EnrolledClassMiddleware), new(*middleware.EnrolledClassMiddlewareImpl)))

var handlerSet = wire.NewSet(
	studentSet,
	classSet,
	enrolledClassSet, wire.Bind(new(controller.StudentController), new(*controller.StudentControllerImpl)), wire.Bind(new(class3.ClassController), new(*class3.ClassControllerImpl)), wire.Bind(new(enrolled_class3.EnrolledClassController), new(*enrolled_class3.EnrolledClassControllerImpl)), wire.Struct(new(app.RouterHandler), "*"), app.NewRouter, middleware.NewAuthMiddleware,
)
