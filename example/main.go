package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spektroskop/keiro"
)

func hello(w http.ResponseWriter, r *http.Request) {
	v := keiro.Context(r, "hei").(string)
	fmt.Fprintf(w, "Hello %v %v\n", keiro.Param(r, "param"), v)
}

func main() {
	ctx := context.WithValue(context.Background(), "hei", "ja")
	mux := keiro.New(ctx)
	mux.GET("/hello/:param", http.HandlerFunc(hello))
	http.ListenAndServe(":3000", mux)
}
