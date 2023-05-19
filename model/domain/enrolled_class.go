package domain

import "github.com/hisyamsk/university-classes-rest-api/entity"

type EnrolledClassByClassId struct {
	ClassId  int
	Students []*entity.Student
}

type EnrolledClassByStudentId struct {
	StudentId int
	Classes   []*entity.Class
}
