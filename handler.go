package keiro

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		code := http.StatusInternalServerError
		if httpError, ok := err.(*HTTPError); ok {
			code = httpError.Code()
		}
		http.Error(w, err.Error(), code)
	}
}
