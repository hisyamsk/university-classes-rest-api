package helper

import (
	"encoding/json"
	"net/http"

	"github.com/hisyamsk/university-classes-rest-api/model/web"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, webResponse *web.WebResponse) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}
