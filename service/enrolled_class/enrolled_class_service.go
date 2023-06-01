package enrolled_class

import (
	"context"

	"github.com/hisyamsk/university-classes-rest-api/model/web/enrolled_class"
)

type EnrolledClassService interface {
	Create(ctx context.Context, req *enrolled_class.EnrolledClassRequest)
	Delete(ctx context.Context, req *enrolled_class.EnrolledClassRequest)
	Find(ctx context.Context, req *enrolled_class.EnrolledClassRequest)
}
