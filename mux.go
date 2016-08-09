package keiro

import (
	"context"
	"net/http"

	"github.com/naoina/denco"
)

type key int

const (
	paramsKey key = iota
)

type Mux struct {
	records map[string][]denco.Record
	router  map[string]*denco.Router
}

func New() *Mux {
	return &Mux{
		records: make(map[string][]denco.Record),
		router:  make(map[string]*denco.Router),
	}
}

func (m *Mux) Handle(method, path string, handler http.Handler) {
	records := m.records[method]
	records = append(records, denco.NewRecord(path, handler))
	m.records[method] = records
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := m.router[r.Method]

	data, params, ok := router.Lookup(r.URL.Path)
	if ok {
		handler := data.(http.Handler)
		ctx := context.WithValue(r.Context(), paramsKey, params)
		handler.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	http.Error(w,
		http.StatusText(http.StatusNotFound),
		http.StatusNotFound,
	)
}

func (m *Mux) Run(address string) error {
	for method, records := range m.records {
		router := denco.New()
		router.Build(records)
		m.router[method] = router
	}

	return http.ListenAndServe(address, m)
}
