package middleware

import (
	"net/http"
	"os"

	"github.com/hisyamsk/university-classes-rest-api/helper"
	"github.com/hisyamsk/university-classes-rest-api/model/web"
	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(router *httprouter.Router) http.Handler {
	return &AuthMiddleware{
		Handler: router,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if os.Getenv("API_KEY_SECRET") == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	webResponse := &web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: http.StatusText(http.StatusUnauthorized),
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusUnauthorized)
}
