package class

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
)

type ClassRepository interface {
	Save(ctx context.Context, tx *sql.Tx, class *entity.Class) *entity.Class
	Update(ctx context.Context, tx *sql.Tx, class *entity.Class) *entity.Class
	Delete(ctx context.Context, tx *sql.Tx, classId int)
	FindById(ctx context.Context, tx *sql.Tx, classId int) (*entity.Class, error)
	FindAll(ctx context.Context, tx *sql.Tx, class *entity.Class) []*entity.Class
}
