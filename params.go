package keiro

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type contextKey string

func (key contextKey) String() string {
	return "keiro/" + string(key)
}

const ParamKey = contextKey("Param")

func Param(r *http.Request, name string) string {
	params := r.Context().Value(ParamKey)
	return params.(httprouter.Params).ByName(name)
}

func makeParams(ctx context.Context, params httprouter.Params) context.Context {
	return context.WithValue(ctx, ParamKey, params)
}
