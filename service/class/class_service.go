package class

import (
	"context"

	webClass "github.com/hisyamsk/university-classes-rest-api/model/web/class"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
)

type ClassService interface {
	Create(ctx context.Context, req *webClass.ClassCreateRequest) *webClass.ClassResponse
	Update(ctx context.Context, req *webClass.ClassUpdateRequest) *webClass.ClassResponse
	Delete(ctx context.Context, classId int)
	FindById(ctx context.Context, classId int) *webClass.ClassResponse
	FindAll(ctx context.Context) []*webClass.ClassResponse
	FindStudentsById(ctx context.Context, classId int) []*webStudent.StudentResponse
}
