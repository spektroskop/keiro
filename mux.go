package keiro

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Mux struct {
	router *httprouter.Router
}

func New() *Mux {
	return &Mux{router: httprouter.New()}
}

func (mux *Mux) Handle(method, path string, handler http.Handler) {
	mux.router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := params(r.Context(), p)
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux.router.ServeHTTP(w, r)
}
