package helper

import (
	"github.com/hisyamsk/university-classes-rest-api/entity"
	"github.com/hisyamsk/university-classes-rest-api/model/web/student"
)

func ToStudentResponse(entity *entity.Student) *student.StudentResponse {
	return &student.StudentResponse{
		Id:       entity.Id,
		Name:     entity.Name,
		Email:    entity.Email,
		Active:   entity.Active,
		Semester: entity.Semester,
	}
}

func ToStudentsResponse(studentsEntity []*entity.Student) []*student.StudentResponse {
	students := []*student.StudentResponse{}
	for _, val := range studentsEntity {
		students = append(students, ToStudentResponse(val))
	}

	return students
}
