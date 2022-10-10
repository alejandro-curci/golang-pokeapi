package errors

import (
	"fmt"
	"net/http"
)

var (
	ErrNotFound   = New("not_found", "resource not found", http.StatusNotFound)
	ErrRestClient = New("rest_client_error", "rest client error", http.StatusInternalServerError)
	ErrStorage    = New("storage_error", "storage  error", http.StatusInternalServerError)
	ErrUnhandled  = New("unhandled_error", "unhandled error", http.StatusInternalServerError)
	ErrBadRequest = New("bad_request", "bad request", http.StatusBadRequest)
	ErrMarshal    = New("marshal_error", "marshal error", http.StatusInternalServerError)
)

type ApiError struct {
	code       string
	msg        string
	httpStatus int
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[code: %s][msg: %s]", e.code, e.msg)
}

func (e ApiError) Code() string {
	return e.code
}

func (e *ApiError) Status() int {
	return e.httpStatus
}

func New(code, msg string, httpCode int) ApiError {
	return ApiError{
		code:       code,
		msg:        msg,
		httpStatus: httpCode,
	}
}
