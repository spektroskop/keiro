package keiro

import "net/http"

type HTTPError struct {
	code    int
	message string
}

func (err *HTTPError) Error() string {
	return err.message
}

func (err *HTTPError) Code() int {
	return err.code
}

func NotFound(message string) error {
	return &HTTPError{
		http.StatusNotFound, message,
	}
}
