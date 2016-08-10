package keiro

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type key int

const paramsKey key = iota

func Param(r *http.Request, name string) string {
	params := r.Context().Value(paramsKey)
	return params.(httprouter.Params).ByName(name)
}

func params(ctx context.Context, params httprouter.Params) context.Context {
	return context.WithValue(ctx, paramsKey, params)
}
