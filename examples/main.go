package main

import (
	"fmt"
	"net/http"

	"github.com/spektroskop/keiro"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v\n", keiro.Param(r, "param"))
}

func main() {
	mux := keiro.New()
	mux.GET("/hello/:param", http.HandlerFunc(hello))
	http.ListenAndServe(":3000", mux)
}
