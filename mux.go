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

func WithValue(req *http.Request, key, object interface{}) *http.Request {
	return req.WithContext(
		context.WithValue(
			req.Context(), key, object,
		),
	)
}

func Value(req *http.Request, key interface{}) interface{} {
	return req.Context().Value(key)
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

func Compose(handlers ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		for i := len(handlers) - 1; i >= 0; i-- {
			handler = handlers[i](handler)
		}

		return handler
	}
}

func Into(hnd http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	return Compose(middleware...)(hnd)
}

func EmptyFunc(http.ResponseWriter, *http.Request) {

}

var Empty = http.HandlerFunc(EmptyFunc)
