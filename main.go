package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/router"
	"api/src/settings"
)

func main() {
	settings.Load()

	r := router.Generate()

	fmt.Printf("Listen on port %d", settings.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", settings.Port), r))
}
