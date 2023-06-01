package enrolled_class

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type EnrolledClassController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
