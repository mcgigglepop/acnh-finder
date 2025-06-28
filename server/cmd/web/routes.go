package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mcgigglepop/acnh-finder/server/internal/config"
	"github.com/mcgigglepop/acnh-finder/server/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/register", handlers.Repo.RegisterGet)
	mux.Post("/register", handlers.Repo.RegisterPost)
	
	mux.Get("/login", handlers.Repo.LoginGet)
	mux.Post("/login", handlers.Repo.LoginPost)

	mux.Get("/email-verification", handlers.Repo.EmailVerificationGet)
	mux.Post("/email-verification", handlers.Repo.EmailVerificationPost)	
	
	mux.Route("/", func(mux chi.Router) {
		mux.Use(Auth) // if you want to apply auth just for these
		mux.Get("/dashboard", handlers.Repo.DashboardGet)
		mux.Get("/choose-hemisphere", handlers.Repo.ChooseHemisphereGet)
		mux.Post("/choose-hemisphere", handlers.Repo.ChooseHemispherePost)
	})

	mux.Route("/fish", func(mux chi.Router) {
		mux.Use(Auth) // if you want to apply auth just for these
		mux.Get("/dashboard", handlers.Repo.FishDashboardGet)
		// New JSON API to fetch filtered fish
		mux.Get("/available", handlers.Repo.GetAvailableFish)

		// single endpoint to handle insert/delete
		mux.Post("/userfish", handlers.Repo.UpdateUserFish)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
