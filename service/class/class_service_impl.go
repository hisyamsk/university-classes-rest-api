package class

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/exception"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	webClass "github.com/hisyamsk/university-classes-rest-api/model/web/class"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	classRepository "github.com/hisyamsk/university-classes-rest-api/repository/class"
)

type ClassServiceImpl struct {
	DB              *sql.DB
	ClassRepository classRepository.ClassRepository
	Validate        *validator.Validate
}

func NewClassService(db *sql.DB, repository classRepository.ClassRepository, validator *validator.Validate) *ClassServiceImpl {
	return &ClassServiceImpl{
		DB:              db,
		ClassRepository: repository,
		Validate:        validator,
	}
}

func (service *ClassServiceImpl) Create(ctx context.Context, req *webClass.ClassCreateRequest) *webClass.ClassResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	class := &entity.Class{Name: req.Name, StartAt: req.StartAt, EndAt: req.EndAt}
	class = service.ClassRepository.Save(ctx, tx, class)

	return helper.ToClassResponse(class)
}

func (service *ClassServiceImpl) Update(ctx context.Context, req *webClass.ClassUpdateRequest) *webClass.ClassResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, findErr := service.ClassRepository.FindById(ctx, tx, req.Id)
	if findErr != nil {
		panic(exception.NewNotFoundError(findErr.Error()))
	}

	classEntity := &entity.Class{Id: req.Id, Name: req.Name, StartAt: req.StartAt, EndAt: req.EndAt}
	class := service.ClassRepository.Update(ctx, tx, classEntity)

	return helper.ToClassResponse(class)
}

func (service *ClassServiceImpl) Delete(ctx context.Context, classId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, findErr := service.ClassRepository.FindById(ctx, tx, classId)
	if findErr != nil {
		panic(exception.NewNotFoundError(findErr.Error()))
	}

	service.ClassRepository.Delete(ctx, tx, classId)
}

func (service *ClassServiceImpl) FindById(ctx context.Context, classId int) *webClass.ClassResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	class, findErr := service.ClassRepository.FindById(ctx, tx, classId)
	if findErr != nil {
		panic(exception.NewNotFoundError(findErr.Error()))
	}

	return helper.ToClassResponse(class)
}

func (service *ClassServiceImpl) FindAll(ctx context.Context) []*webClass.ClassResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	classes := service.ClassRepository.FindAll(ctx, tx)

	return helper.ToClassesResponse(classes)
}

func (service *ClassServiceImpl) FindStudentsById(ctx context.Context, classId int) []*webStudent.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, findErr := service.ClassRepository.FindById(ctx, tx, classId)
	if findErr != nil {
		panic(exception.NewNotFoundError(findErr.Error()))
	}
	enrolledStudents := service.ClassRepository.FindStudentsById(ctx, tx, classId)

	return helper.ToStudentsResponse(enrolledStudents)
}
