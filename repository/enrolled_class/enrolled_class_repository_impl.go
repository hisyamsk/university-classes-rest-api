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

func (repository *EnrolledClassRepositoryImpl) FindByClassId(ctx context.Context, tx *sql.Tx, classId int) []*entity.Student {
	query := "SELECT student.id, student.name, student.email, student.active, student.semester FROM enrolled_class JOIN student ON student_id = student.id WHERE class_id = $1"
	rows, err := tx.QueryContext(ctx, query, classId)
	helper.PanicIfError(err)
	defer rows.Close()

	enrolledStudents := []*entity.Student{}

	for rows.Next() {
		student := &entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Active, &student.Semester)
		helper.PanicIfError(err)
		enrolledStudents = append(enrolledStudents, student)
	}

	return enrolledStudents
}

func (repository *EnrolledClassRepositoryImpl) FindByStudentId(ctx context.Context, tx *sql.Tx, studentId int) []*entity.Class {
	query := "SELECT class_id, class.name, class.start_at, class.end_at FROM enrolled_class JOIN class ON class_id = class.id WHERE student_id = $1"
	rows, err := tx.QueryContext(ctx, query, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	enrolledClasses := []*entity.Class{}

	for rows.Next() {
		class := &entity.Class{}
		err := rows.Scan(&class.Id, &class.Name, &class.StartAt, &class.EndAt)
		helper.PanicIfError(err)
		enrolledClasses = append(enrolledClasses, class)
	}

	return enrolledClasses
}
