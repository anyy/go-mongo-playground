package infrastructure

import (
	"github.com/gazelle0130/go-mongo-playground/src/app/interfaces/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var Router *chi.Mux

func init() {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	kvsh, err := NewKVSHandler()
	if err != nil {
		panic(err.Error())
	}
	uc := controllers.NewUserController(kvsh)
	r.Route("/user", func(r chi.Router) {
		r.Post("/", uc.Create)
		r.Get("/", uc.Index)
		r.Route("/{userID}", func(r chi.Router) {
			r.Delete("/", uc.Delete)
		})
	})

	Router = r
}
