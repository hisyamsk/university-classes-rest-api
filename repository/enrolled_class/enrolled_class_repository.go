package enrolled_class

import (
	"context"
	"database/sql"

	"github.com/hisyamsk/university-classes-rest-api/entity"
)

type EnrolledClassRepository interface {
	Save(ctx context.Context, tx *sql.Tx, enrolledClass *entity.EnrolledClass) *entity.EnrolledClass
}
