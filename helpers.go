package keiro

import (
	"context"

	"github.com/naoina/denco"
)

func Param(ctx context.Context, name string) string {
	params := ctx.Value(paramsKey).(denco.Params)
	return params.Get(name)
}
