package app

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, apiCfg ApiConfig) {
	router.Delete("/url/{id}", apiCfg.DeleteURL)
	router.Post("/shorten", apiCfg.ShortenURL)
	router.Get("/url", apiCfg.GetURLs)
	router.Get("/url/{id}", apiCfg.GetURLByID)
	router.Get("/original", apiCfg.GetURLByOriginalURL)
	router.Get("/{shortURL}", apiCfg.RedirectURL)
	router.Put("/url/{id}", apiCfg.UpdateURL)
}
