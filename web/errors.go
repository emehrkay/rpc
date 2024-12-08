package web

import (
	"fmt"
	"net/http"
)

var (
	InvalidUUIDErr = HttpError{
		Message:    "invalid uuid",
		StatusCode: http.StatusBadRequest,
	}

	RequestBodyRequiredErr = HttpError{
		Message:    "request body required",
		StatusCode: http.StatusBadRequest,
	}

	RequestBodyInvalidErr = HttpError{
		Message:    "request body invalid",
		StatusCode: http.StatusBadRequest,
	}
)

type WebError struct {
	Message string `json:"message"`
}

type HttpError struct {
	OrignialError error
	Message       string
	StatusCode    int
}

func (h HttpError) Error() string {
	err := fmt.Sprintf(`[%d] %v`, h.StatusCode, h.Message)
	if h.OrignialError != nil {
		err = err + " -- " + h.OrignialError.Error()
	}

	return err
}

func (h HttpError) WebError() WebError {
	if h.Message != "" {
		return WebError{Message: h.Message}
	}

	if h.StatusCode > 0 {
		return WebError{Message: http.StatusText(h.StatusCode)}
	}

	return WebError{Message: "error"}
}
