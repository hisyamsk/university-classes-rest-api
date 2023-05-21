package student

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	webClass "github.com/hisyamsk/university-classes-rest-api/model/web/class"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	"github.com/hisyamsk/university-classes-rest-api/repository/student"
)

type StudentServiceImpl struct {
	DB                *sql.DB
	StudentRepository student.StudentRepository
}

func NewStudentService(repository student.StudentRepository, db *sql.DB) *StudentServiceImpl {
	return &StudentServiceImpl{
		StudentRepository: repository,
		DB:                db,
	}
}

func (service *StudentServiceImpl) Create(ctx context.Context, req *webStudent.StudentCreateRequest) *webStudent.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student := &entity.Student{Name: req.Name, Email: req.Email, Active: req.Active, Semester: req.Semester}
	student = service.StudentRepository.Save(ctx, tx, student)

	return helper.ToStudentResponse(student)
}
func (service *StudentServiceImpl) Update(ctx context.Context, req *webStudent.StudentUpdateRequest) *webStudent.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student := &entity.Student{Id: req.Id, Name: req.Name, Email: req.Email, Active: req.Active, Semester: req.Semester}
	student = service.StudentRepository.Update(ctx, tx, student)

	return helper.ToStudentResponse(student)
}
func (service *StudentServiceImpl) Delete(ctx context.Context, studentId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.StudentRepository.Delete(ctx, tx, studentId)
}
func (service *StudentServiceImpl) FindById(ctx context.Context, studentId int) (*webStudent.StudentResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		return nil, err
	}

	return helper.ToStudentResponse(student), nil
}
func (service *StudentServiceImpl) FindAll(ctx context.Context) []*webStudent.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students := service.StudentRepository.FindAll(ctx, tx)
	return helper.ToStudentsResponse(students)
}

func (service *StudentServiceImpl) FindClasses(ctx context.Context, studentId int) []*webClass.ClassResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	enrolledClasses := service.StudentRepository.FindClassesById(ctx, tx, studentId)
	return helper.ToClassesResponse(enrolledClasses)
}
