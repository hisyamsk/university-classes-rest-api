package enrolled_class

import (
	"context"
	"database/sql"
	"fmt"

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

func (repository *EnrolledClassRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, enrolledClass *entity.EnrolledClass) {
	query := "DELETE FROM enrolled_class WHERE student_id = $1 AND class_id = $2 RETURNING id"
	row := tx.QueryRowContext(ctx, query, enrolledClass.StudentId, enrolledClass.ClassId)
	err := row.Scan(&enrolledClass.Id)
	helper.PanicIfError(err)
}

func (repository *EnrolledClassRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, enrolledClass *entity.EnrolledClass) (*entity.EnrolledClass, error) {
	query := "SELECT id, student_id, class_id FROM enrolled_class WHERE student_id = $1 AND class_id = $2"
	rows, err := tx.QueryContext(ctx, query, enrolledClass.StudentId, enrolledClass.StudentId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return enrolledClass, nil
	}

	return nil, fmt.Errorf("studentId or classId was not found!")
}
