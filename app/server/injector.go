//go:build wireinject
// +build wireinject

package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/db"
	classController "github.com/hisyamsk/university-classes-rest-api/controller/class"
	enrolledClassController "github.com/hisyamsk/university-classes-rest-api/controller/enrolled_class"
	studentController "github.com/hisyamsk/university-classes-rest-api/controller/student"
	"github.com/hisyamsk/university-classes-rest-api/middleware"
	classRepository "github.com/hisyamsk/university-classes-rest-api/repository/class"
	enrolledClassRepository "github.com/hisyamsk/university-classes-rest-api/repository/enrolled_class"
	studentRepository "github.com/hisyamsk/university-classes-rest-api/repository/student"
	classService "github.com/hisyamsk/university-classes-rest-api/service/class"
	enrolledClassService "github.com/hisyamsk/university-classes-rest-api/service/enrolled_class"
	studentService "github.com/hisyamsk/university-classes-rest-api/service/student"
)

var classSet = wire.NewSet(
	classRepository.NewClassRepositoryImpl,
	wire.Bind(new(classRepository.ClassRepository), new(*classRepository.ClassRepositoryImpl)),
	classService.NewClassService,
	wire.Bind(new(classService.ClassService), new(*classService.ClassServiceImpl)),
	classController.NewClassController,
)

var studentSet = wire.NewSet(
	studentRepository.NewStudentRepository,
	wire.Bind(new(studentRepository.StudentRepository), new(*studentRepository.StudentRepositoryImpl)),
	studentService.NewStudentService,
	wire.Bind(new(studentService.StudentService), new(*studentService.StudentServiceImpl)),
	studentController.NewStudentController,
)

var enrolledClassSet = wire.NewSet(
	enrolledClassRepository.NewEnrolledClassRepository,
	wire.Bind(new(enrolledClassRepository.EnrolledClassRepository), new(*enrolledClassRepository.EnrolledClassRepositoryImpl)),
	enrolledClassService.NewEnrolledClassService,
	wire.Bind(new(enrolledClassService.EnrolledClassService), new(*enrolledClassService.EnrolledClassServiceImpl)),
	enrolledClassController.NewEnrolledClassController,
	middleware.NewEnrolledClassMiddleware,
	wire.Bind(new(middleware.EnrolledClassMiddleware), new(*middleware.EnrolledClassMiddlewareImpl)),
)

var handlerSet = wire.NewSet(
	studentSet,
	classSet,
	enrolledClassSet,
	wire.Bind(new(studentController.StudentController), new(*studentController.StudentControllerImpl)),
	wire.Bind(new(classController.ClassController), new(*classController.ClassControllerImpl)),
	wire.Bind(new(enrolledClassController.EnrolledClassController), new(*enrolledClassController.EnrolledClassControllerImpl)),
	wire.Struct(new(app.RouterHandler), "*"),
	app.NewRouter,
	middleware.NewAuthMiddleware,
)

func InitializeHandler(dbName string) http.Handler {
	wire.Build(
		db.NewDBConnection,
		validator.New,
		handlerSet,
	)
	return nil
}
