package enrolledclass

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
)

type EnrolledClassRepository interface {
	FindByClassId(ctx context.Context, tx *sql.Tx, enrolledClassId int) *entity.EnrolledClass
	FindByStudentId(ctx context.Context, tx *sql.Tx, enrolledClassId int) []*entity.Class
}
