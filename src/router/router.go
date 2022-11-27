package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

//Return a router with routers config
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.Config(r)
}
