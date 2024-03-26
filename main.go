package main

import (
	"context"
	"net/http"

	"github.com/brendenehlers/oapi-gen-test/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)



func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	api := handler.New(context.Background())
	r.Mount("/", api.Handler())

	http.ListenAndServe(":8080", r)
}