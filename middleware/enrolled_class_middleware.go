package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/web/enrolled_class"
	"github.com/hisyamsk/university-classes-rest-api/service/class"
	"github.com/hisyamsk/university-classes-rest-api/service/student"
	"github.com/julienschmidt/httprouter"
)

type EnrolledClassMiddleware interface {
	Guard(next httprouter.Handle) httprouter.Handle
}

type EnrolledClassMiddlewareImpl struct {
	classService   class.ClassService
	studentService student.StudentService
}

func NewEnrolledClassMiddleware(classService class.ClassService, studentService student.StudentService) *EnrolledClassMiddlewareImpl {
	return &EnrolledClassMiddlewareImpl{
		classService:   classService,
		studentService: studentService,
	}
}

func (middleware *EnrolledClassMiddlewareImpl) Guard(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		bodyByte, _ := io.ReadAll(r.Body)
		requestBody := &enrolled_class.EnrolledClassRequest{}
		err := json.Unmarshal(bodyByte, requestBody)
		helper.PanicIfError(err)

		middleware.studentService.FindById(r.Context(), requestBody.StudentId)
		middleware.classService.FindById(r.Context(), requestBody.ClassId)

		r.Body = io.NopCloser(bytes.NewBuffer(bodyByte))

		next(w, r, p)
	}
}
