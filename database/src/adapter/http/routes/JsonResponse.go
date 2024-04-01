package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Writer     http.ResponseWriter `json:"-"`
	Key        string              `json:"key"`
	Err        string              `json:"error"`
	StatusCode int                 `json:"statusCode"`
}

func NewJsonResponse(
	writer http.ResponseWriter,
	key string,
	err string,
	statusCode int,
) *JsonResponse {
	return &JsonResponse{
		Writer:     writer,
		Key:        key,
		Err:        err,
		StatusCode: statusCode,
	}
}

func (jsonResponse *JsonResponse) ThrowError() {
	jsonResponse.Writer.Header().Set("Content-type", "application/json")
	jsonResponse.Writer.WriteHeader(jsonResponse.StatusCode)

	responseMarshal, err := json.Marshal(jsonResponse)

	if err != nil {
		log.Fatal(err)
	}

	jsonResponse.Writer.Write(responseMarshal)
}
