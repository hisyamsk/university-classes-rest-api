package student

import (
	"context"

	webClass "github.com/hisyamsk/university-classes-rest-api/model/web/class"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
)

type StudentService interface {
	Create(ctx context.Context, req *webStudent.StudentCreateRequest) *webStudent.StudentResponse
	Update(ctx context.Context, req *webStudent.StudentUpdateRequest) *webStudent.StudentResponse
	Delete(ctx context.Context, studentId int)
	FindById(ctx context.Context, studentId int) (*webStudent.StudentResponse, error)
	FindAll(ctx context.Context) []*webStudent.StudentResponse
	FindClasses(ctx context.Context, studentId int) []*webClass.ClassResponse
}
