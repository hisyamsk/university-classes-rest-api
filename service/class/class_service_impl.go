package class

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	webClass "github.com/hisyamsk/university-classes-rest-api/model/web/class"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	classRepository "github.com/hisyamsk/university-classes-rest-api/repository/class"
)

type ClassServiceImpl struct {
	DB              *sql.DB
	ClassRepository classRepository.ClassRepository
}

func NewClassService(db *sql.DB, repository classRepository.ClassRepository) *ClassServiceImpl {
	return &ClassServiceImpl{
		DB:              db,
		ClassRepository: repository,
	}
}

func (service *ClassServiceImpl) Create(ctx context.Context, req *webClass.ClassCreateRequest) *webClass.ClassResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	class := &entity.Class{Name: req.Name, StartAt: req.StartAt, EndAt: req.EndAt}
	class = service.ClassRepository.Save(ctx, tx, class)

	return helper.ToClassResponse(class)
}

func (service *ClassServiceImpl) Update(ctx context.Context, req *webClass.ClassUpdateRequest) *webClass.ClassResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, errClass := service.ClassRepository.FindById(ctx, tx, req.Id)
	if errClass != nil {
		helper.PanicIfError(err)
	}

	classEntity := &entity.Class{Id: req.Id, Name: req.Name, StartAt: req.StartAt, EndAt: req.EndAt}
	class := service.ClassRepository.Update(ctx, tx, classEntity)

	return helper.ToClassResponse(class)
}

func (service *ClassServiceImpl) Delete(ctx context.Context, classId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, errClass := service.ClassRepository.FindById(ctx, tx, classId)
	if errClass != nil {
		helper.PanicIfError(err)
	}

	service.ClassRepository.Delete(ctx, tx, classId)
}

func (service *ClassServiceImpl) FindById(ctx context.Context, classId int) (*webClass.ClassResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	class, err := service.ClassRepository.FindById(ctx, tx, classId)
	if err != nil {
		return nil, fmt.Errorf("Class with id:%d was not found", classId)
	}

	return helper.ToClassResponse(class), nil
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

	enrolledStudents := service.ClassRepository.FindStudentsById(ctx, tx, classId)

	return helper.ToStudentsResponse(enrolledStudents)
}
