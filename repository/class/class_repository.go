package repository

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/model/domain"
)

type ClassRepository interface {
	Save(ctx context.Context, tx *sql.Tx, class *domain.Class) *domain.Class
	Update(ctx context.Context, tx *sql.Tx, class *domain.Class) *domain.Class
	Delete(ctx context.Context, tx *sql.Tx, classId int)
	FindById(ctx context.Context, tx *sql.Tx, classId int) (*domain.Class, error)
	FindAll(ctx context.Context, tx *sql.Tx, class *domain.Class) []*domain.Class
}