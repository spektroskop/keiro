package keiro

import "net/http"

func (mux *Mux) GET(path string, handler http.Handler) {
	mux.Handle("GET", path, handler)
}

func (mux *Mux) POST(path string, handler http.Handler) {
	mux.Handle("POST", path, handler)
}

func (mux *Mux) PUT(path string, handler http.Handler) {
	mux.Handle("PUT", path, handler)
}

func (mux *Mux) HEAD(path string, handler http.Handler) {
	mux.Handle("HEAD", path, handler)
}

func (mux *Mux) DELETE(path string, handler http.Handler) {
	mux.Handle("DELETE", path, handler)
}

func (mux *Mux) PATCH(path string, handler http.Handler) {
	mux.Handle("PATCH", path, handler)
}
