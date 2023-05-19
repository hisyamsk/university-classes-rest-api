package enrolledclass

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
)

type EnrolledClassRepositoryImpl struct {
}

func NewEnrolledClassRepository() *EnrolledClassRepositoryImpl {
	return &EnrolledClassRepositoryImpl{}
}

func (repository *EnrolledClassRepositoryImpl) FindByClassId(ctx context.Context, tx *sql.Tx, classId int) *entity.EnrolledClass {
	query := "SELECT id, student.id, student.name, FROM enrolled_class JOIN student ON student_id = student.id WHERE class_id = ?"
	rows, err := tx.QueryContext(ctx, query, classId)
	helper.PanicIfError(err)
	defer rows.Close()

	enrolledClass := &entity.EnrolledClass{
		ClassId:  classId,
		Students: []*entity.Student{},
	}

	for rows.Next() {
		var ignoreValue any
		student := &entity.Student{}
		err := rows.Scan(ignoreValue, &student.Id, &student.Name)
		helper.PanicIfError(err)
		enrolledClass.Students = append(enrolledClass.Students, student)
	}

	return enrolledClass
}

func (repository *EnrolledClassRepositoryImpl) FindByStudentId(ctx context.Context, tx *sql.Tx, studentId int) []*entity.Class {
	panic("not implemented") // TODO: Implement
}
