package enrolled_class

import (
	"net/http"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/web"
	"github.com/hisyamsk/university-classes-rest-api/model/web/enrolled_class"
	enrolledClassService "github.com/hisyamsk/university-classes-rest-api/service/enrolled_class"
	"github.com/julienschmidt/httprouter"
)

type EnrolledClassControllerImpl struct {
	Service enrolledClassService.EnrolledClassService
}

func NewEnrolledClassController(service enrolledClassService.EnrolledClassService) *EnrolledClassControllerImpl {
	return &EnrolledClassControllerImpl{
		Service: service,
	}
}

func (controller *EnrolledClassControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	enrolledClassRequest := &enrolled_class.EnrolledClassRequest{}
	helper.ReadFromRequestBody(request, enrolledClassRequest)

	controller.Service.Create(request.Context(), enrolledClassRequest)
	webResponse := &web.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusCreated)
}

func (controller *EnrolledClassControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	enrolledClassRequest := &enrolled_class.EnrolledClassRequest{}
	helper.ReadFromRequestBody(request, enrolledClassRequest)

	controller.Service.Delete(request.Context(), enrolledClassRequest)
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}
