package router

import "github.com/gorilla/mux"

//Return a router with routers config
func Generate() *mux.Router {
	return mux.NewRouter()
}
