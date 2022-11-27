package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

func Config(r *mux.Router) *mux.Router {
	routers := userRouters

	for _, router := range routers {
		r.HandleFunc(router.Uri, router.Function).Methods(router.Method)
	}

	return r
}
