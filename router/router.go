package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type LocaleAPI struct {
	Router *chi.Mux
}

func NewAPI() (*LocaleAPI, error) {
	var api LocaleAPI

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(EnsureValidToken())
	r.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to locale-mgmt api"))
	})

	api.Router = r
	return &api, nil

}
