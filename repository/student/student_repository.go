package repository

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
)

type StudentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.Student
	Update(ctx context.Context, tx *sql.Tx, student *entity.Student) *entity.Student
	Delete(ctx context.Context, tx *sql.Tx, studentId int)
	FindById(ctx context.Context, tx *sql.Tx, studentId int) (*entity.Student, error)
	FindAll(ctx context.Context, tx *sql.Tx, student *entity.Student) []*entity.Student
}
