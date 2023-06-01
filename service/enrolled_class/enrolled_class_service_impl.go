package enrolled_class

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/exception"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/web/enrolled_class"
	enrolledClassRepository "github.com/hisyamsk/university-classes-rest-api/repository/enrolled_class"
)

type EnrolledClassServiceImpl struct {
	DB         *sql.DB
	Repository enrolledClassRepository.EnrolledClassRepository
	Validate   *validator.Validate
}

func NewEnrolledClassService(db *sql.DB, repository enrolledClassRepository.EnrolledClassRepository, validator *validator.Validate) *EnrolledClassServiceImpl {
	return &EnrolledClassServiceImpl{
		DB:         db,
		Repository: repository,
		Validate:   validator,
	}
}

func (service *EnrolledClassServiceImpl) Create(ctx context.Context, req *enrolled_class.EnrolledClassRequest) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	enrolledClassEntity := &entity.EnrolledClass{StudentId: req.StudentId, ClassId: req.ClassId}

	service.Repository.Save(ctx, tx, enrolledClassEntity)
}

func (service *EnrolledClassServiceImpl) Delete(ctx context.Context, req *enrolled_class.EnrolledClassRequest) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	enrolledClassEntity := &entity.EnrolledClass{StudentId: req.StudentId, ClassId: req.ClassId}
	_, findErr := service.Repository.Find(ctx, tx, enrolledClassEntity)
	if findErr != nil {
		panic(exception.NewNotFoundError(findErr.Error()))
	}

	service.Repository.Delete(ctx, tx, enrolledClassEntity)
}

func (service *EnrolledClassServiceImpl) Find(ctx context.Context, req *enrolled_class.EnrolledClassRequest) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	enrolledClassEntity := &entity.EnrolledClass{StudentId: req.StudentId, ClassId: req.ClassId}
	_, findErr := service.Repository.Find(ctx, tx, enrolledClassEntity)
	if findErr != nil {
		panic(exception.NewNotFoundError(findErr.Error()))
	}
}
