package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/spektroskop/keiro"
)

func foo(w http.ResponseWriter, r *http.Request) error {
	param := keiro.Param(r.Context(), "param")
	logrus.Infof("Foo %v", param)
	return keiro.JSON(w,
		http.StatusOK,
		map[string]interface{}{"Foo": "OK"},
	)
}

func bar(w http.ResponseWriter, r *http.Request) {
	param := keiro.Param(r.Context(), "param")
	logrus.Infof("Bar %v", param)
	keiro.JSON(w,
		http.StatusInternalServerError,
		map[string]interface{}{"Bar": "Error"},
	)
}

func baz(w http.ResponseWriter, r *http.Request) error {
	param := keiro.Param(r.Context(), "param")
	logrus.Infof("Baz %v", param)
	return keiro.Forbidden("Nope")
}

func main() {
	mux := keiro.New()

	mux.GET("/foo/:param", keiro.Handler(foo))
	mux.GET("/bar/:param", http.HandlerFunc(bar))
	mux.GET("/baz", keiro.Handler(baz))

	logrus.Fatal(mux.Run(":3000"))
}
