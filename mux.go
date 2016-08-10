package keiro

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Mux struct {
	ctx    context.Context
	router *httprouter.Router
}

func Context(r *http.Request, name string) interface{} {
	return r.Context().Value(name)
}

func New(ctx context.Context) *Mux {
	return &Mux{ctx, httprouter.New()}
}

func (mux *Mux) Handle(method, path string, handler http.Handler) {
	mux.router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		r = r.WithContext(mux.ctx)
		r = r.WithContext(params(r.Context(), p))
		handler.ServeHTTP(w, r)
	})
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux.router.ServeHTTP(w, r)
}
