package main

import (
	"net/http"

	"github.com/brendenehlers/oapi-gen-test/handler"
	"github.com/go-chi/chi/v5"
)



func main() {
	r := chi.NewRouter()
	api := handler.New()
	r.Mount("/", api.Handler())

	http.ListenAndServe(":8080", r)
}