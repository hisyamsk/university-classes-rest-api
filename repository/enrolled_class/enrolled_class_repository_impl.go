package enrolledclass

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/domain"
)

type EnrolledClassRepositoryImpl struct {
}

func NewEnrolledClassRepository() *EnrolledClassRepositoryImpl {
	return &EnrolledClassRepositoryImpl{}
}

func (repository *EnrolledClassRepositoryImpl) FindByClassId(ctx context.Context, tx *sql.Tx, classId int) *domain.EnrolledClassByClassId {
	query := "SELECT student.id, student.name, student.email, student.active, student.semester FROM enrolled_class JOIN student ON student_id = student.id WHERE class_id = $1"
	rows, err := tx.QueryContext(ctx, query, classId)
	helper.PanicIfError(err)
	defer rows.Close()

	enrolledStudents := &domain.EnrolledClassByClassId{
		ClassId:  classId,
		Students: []*entity.Student{},
	}

	for rows.Next() {
		student := &entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Active, &student.Semester)
		helper.PanicIfError(err)
		enrolledStudents.Students = append(enrolledStudents.Students, student)
	}

	return enrolledStudents
}

func (repository *EnrolledClassRepositoryImpl) FindByStudentId(ctx context.Context, tx *sql.Tx, studentId int) *domain.EnrolledClassByStudentId {
	query := "SELECT class_id, class.name, class.start_at, class.end_at FROM enrolled_class JOIN class ON class_id = class.id WHERE student_id = $1"
	rows, err := tx.QueryContext(ctx, query, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	enrolledClasses := &domain.EnrolledClassByStudentId{
		StudentId: studentId,
		Classes:   []*entity.Class{},
	}

	for rows.Next() {
		class := &entity.Class{}
		err := rows.Scan(&class.Id, &class.Name, &class.StartAt, &class.EndAt)
		helper.PanicIfError(err)
		enrolledClasses.Classes = append(enrolledClasses.Classes, class)
	}

	return enrolledClasses
}
