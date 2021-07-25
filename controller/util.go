package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	BadRequest = HttpError{ErrorCode: "400", ErrorDesc: "Bad request"}
)

type HttpError struct {
	ErrorCode string `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

func NewHttpError(errorCode, errorDesc string) HttpError {
	return HttpError{
		ErrorCode: errorCode,
		ErrorDesc: errorDesc,
	}
}

func (e HttpError) Error() string {
	return fmt.Sprintf("errorCode: %v, errorDesc: %v", e.ErrorCode, e.ErrorDesc)
}

func WriteResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)
	w.Write(body)
}

func WriteErrResponse(w http.ResponseWriter, statusCode int, body HttpError) {
	bytes, _ := json.Marshal(body)

	log.Println(body.Error())

	w.WriteHeader(statusCode)
	w.Write(bytes)
}
