package keiro

import "net/http"

func (m *Mux) GET(path string, handler http.Handler) {
	m.Handle("GET", path, handler)
}

func (m *Mux) POST(path string, handler http.Handler) {
	m.Handle("POST", path, handler)
}

func (m *Mux) PUT(path string, handler http.Handler) {
	m.Handle("PUT", path, handler)
}

func (m *Mux) PATCH(path string, handler http.Handler) {
	m.Handle("PATCH", path, handler)
}

func (m *Mux) HEAD(path string, handler http.Handler) {
	m.Handle("HEAD", path, handler)
}
