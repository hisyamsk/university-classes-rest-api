package class

import (
	"net/http"
	"strconv"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/web"
	"github.com/hisyamsk/university-classes-rest-api/model/web/class"
	classService "github.com/hisyamsk/university-classes-rest-api/service/class"
	"github.com/julienschmidt/httprouter"
)

type ClassControllerImpl struct {
	ClassService classService.ClassService
}

func NewClassController(service classService.ClassService) *ClassControllerImpl {
	return &ClassControllerImpl{
		ClassService: service,
	}
}

func (controller *ClassControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	classCreateRequest := &class.ClassCreateRequest{}
	helper.ReadFromRequestBody(request, classCreateRequest)

	classResponse := controller.ClassService.Create(request.Context(), classCreateRequest)
	webResponse := &web.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   classResponse,
	}

	writer.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ClassControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	classUpdateRequest := &class.ClassUpdateRequest{}
	helper.ReadFromRequestBody(request, classUpdateRequest)

	classResponse := controller.ClassService.Update(request.Context(), classUpdateRequest)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   classResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ClassControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	classId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	controller.ClassService.Delete(request.Context(), classId)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ClassControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	classId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	classResponse := controller.ClassService.FindById(request.Context(), classId)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   classResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ClassControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	classResponse := controller.ClassService.FindAll(request.Context())
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   classResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *ClassControllerImpl) GetStudentsById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	classId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	studentsResponse := controller.ClassService.FindStudentsById(request.Context(), classId)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentsResponse,
	}

	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}
