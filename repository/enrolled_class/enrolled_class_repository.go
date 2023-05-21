package enrolledclass

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
)

type EnrolledClassRepository interface {
	Save(ctx context.Context, tx *sql.Tx, enrolledClass *entity.EnrolledClass) *entity.EnrolledClass
	FindByClassId(ctx context.Context, tx *sql.Tx, classId int) []*entity.Student
	FindByStudentId(ctx context.Context, tx *sql.Tx, studentId int) []*entity.Class
}
