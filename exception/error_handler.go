package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")

		webResponse := &web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse, http.StatusBadRequest)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(*NotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")

		webResponse := &web.WebResponse{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse, http.StatusNotFound)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	writer.Header().Add("Content-Type", "application/json")

	webResponse := &web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusInternalServerError)
}
