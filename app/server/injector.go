//go:build wireinject
// +build wireinject

package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/hisyamsk/university-classes-rest-api/app"
	"github.com/hisyamsk/university-classes-rest-api/app/db"
	classController "github.com/hisyamsk/university-classes-rest-api/controller/class"
	studentController "github.com/hisyamsk/university-classes-rest-api/controller/student"
	classRepository "github.com/hisyamsk/university-classes-rest-api/repository/class"
	studentRepository "github.com/hisyamsk/university-classes-rest-api/repository/student"
	classService "github.com/hisyamsk/university-classes-rest-api/service/class"
	studentService "github.com/hisyamsk/university-classes-rest-api/service/student"
	"github.com/julienschmidt/httprouter"
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

var routerSet = wire.NewSet(
	studentSet,
	classSet,
	wire.Bind(new(studentController.StudentController), new(*studentController.StudentControllerImpl)),
	wire.Bind(new(classController.ClassController), new(*classController.ClassControllerImpl)),
	wire.Struct(new(app.RouterHandler), "*"),
	app.NewRouter,
)

func InitializeServer(dbName string) *httprouter.Router {
	wire.Build(
		db.NewDBConnection,
		validator.New,
		routerSet,
	)
	return nil
}
