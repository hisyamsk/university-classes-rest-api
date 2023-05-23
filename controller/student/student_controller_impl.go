package controller

import (
	"net/http"
	"strconv"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	web "github.com/hisyamsk/university-classes-rest-api/model/web"
	webStudent "github.com/hisyamsk/university-classes-rest-api/model/web/student"
	"github.com/hisyamsk/university-classes-rest-api/service/student"
	"github.com/julienschmidt/httprouter"
)

type StudentControllerImpl struct {
	StudentService student.StudentService
}

func NewStudentController(service student.StudentService) *StudentControllerImpl {
	return &StudentControllerImpl{
		StudentService: service,
	}
}

func (controller *StudentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentCreateRequest := &webStudent.StudentCreateRequest{}
	helper.ReadFromRequestBody(request, studentCreateRequest)

	studentResponse := controller.StudentService.Create(request.Context(), studentCreateRequest)
	webResponse := &web.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   studentResponse,
	}

	writer.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentUpdateRequest := &webStudent.StudentUpdateRequest{}
	helper.ReadFromRequestBody(request, studentUpdateRequest)

	studentResponse := controller.StudentService.Update(request.Context(), studentUpdateRequest)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	controller.StudentService.Delete(request.Context(), studentId)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	studentResponse := controller.StudentService.FindById(request.Context(), studentId)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentResponse := controller.StudentService.FindAll(request.Context())
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) GetClassesById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	studentResponse := controller.StudentService.FindClasses(request.Context(), studentId)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}
