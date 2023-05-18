package repository

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/model/domain"
)

type StudentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, student *domain.Student) *domain.Student
	Update(ctx context.Context, tx *sql.Tx, student *domain.Student) *domain.Student
	Delete(ctx context.Context, tx *sql.Tx, studentId int)
	FindById(ctx context.Context, tx *sql.Tx, studentId int) (*domain.Student, error)
	FindAll(ctx context.Context, tx *sql.Tx, student *domain.Student) []*domain.Student
}
