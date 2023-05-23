package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StudentController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetClassesById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
