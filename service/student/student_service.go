package student

import (
	"context"

	web "github.com/hisyamsk/university-classes-rest-api/model/web/student"
)

type StudentService interface {
	Create(ctx context.Context, req *web.StudentCreateRequest) *web.StudentResponse
	Update(ctx context.Context, req *web.StudentUpdateRequest) *web.StudentResponse
	Delete(ctx context.Context, studentId int)
	FindById(ctx context.Context, studentId int) (*web.StudentResponse, error)
	FindAll(ctx context.Context) []*web.StudentResponse
}
