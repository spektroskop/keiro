package keiro

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/naoina/denco"
)

func Param(ctx context.Context, name string) string {
	params := ctx.Value(paramsKey).(denco.Params)
	return params.Get(name)
}

func JSON(w http.ResponseWriter, object interface{}) error {
	return json.NewEncoder(w).Encode(object)
}
