package main

import (
	"log"
	"net/http"

	"github.com/daniel/apiRest2/route"
	"github.com/gorilla/mux"
)

func main() {
	port := ":8000"
	router := mux.NewRouter()

	route.StartRoute(router)

	log.Printf("El servidor esta corriendo en el puerto %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
