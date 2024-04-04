package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Writer     http.ResponseWriter `json:"-"`
	Message    string              `json:"message"`
	StatusCode int                 `json:"statusCode"`
}

func NewJsonResponse(
	writer http.ResponseWriter,
	message string,
	statusCode int,
) *JsonResponse {
	return &JsonResponse{
		Writer:     writer,
		Message:    message,
		StatusCode: statusCode,
	}
}

func (jsonResponse *JsonResponse) SendJson() {
	jsonResponse.Writer.Header().Set("Content-type", "application/json")
	jsonResponse.Writer.WriteHeader(jsonResponse.StatusCode)

	responseMarshal, err := json.Marshal(jsonResponse)

	if err != nil {
		log.Fatal(err)
	}

	jsonResponse.Writer.Write(responseMarshal)
}
