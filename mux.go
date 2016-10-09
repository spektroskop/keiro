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

func Context(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}

func WithContext(ctx context.Context) *Mux {
	return &Mux{ctx, httprouter.New()}
}

func New() *Mux {
	return &Mux{context.Background(), httprouter.New()}
}

func (mux *Mux) Handle(method, path string, handler http.Handler) {
	mux.router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := makeParams(mux.ctx, p)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux.router.ServeHTTP(w, r)
}
