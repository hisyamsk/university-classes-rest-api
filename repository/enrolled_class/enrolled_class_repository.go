package enrolledclass

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/model/domain"
)

type EnrolledClassRepository interface {
	FindByClassId(ctx context.Context, tx *sql.Tx, enrolledClassId int) *domain.EnrolledClassByClassId
	FindByStudentId(ctx context.Context, tx *sql.Tx, enrolledClassId int) *domain.EnrolledClassByStudentId
}
