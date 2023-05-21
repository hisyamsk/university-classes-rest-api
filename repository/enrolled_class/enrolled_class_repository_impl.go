package enrolled_class

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

func (repository *EnrolledClassRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, enrolledClass *entity.EnrolledClass) *entity.EnrolledClass {
	query := "INSERT INTO enrolled_class(student_id, class_id) VALUES ($1, $2) RETURNING id"
	row := tx.QueryRowContext(ctx, query, enrolledClass.StudentId, enrolledClass.ClassId)
	err := row.Scan(&enrolledClass.Id)
	helper.PanicIfError(err)

	return enrolledClass
}
