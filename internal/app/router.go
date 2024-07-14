package app

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, apiCfg ApiConfig) {
	router.HandleFunc("/shorten", apiCfg.ShortenURL).Methods("POST")
	router.HandleFunc("/url", apiCfg.GetURLs).Methods("GET")
	router.HandleFunc("/url/{id}", apiCfg.GetURLByID).Methods("GET")
	router.HandleFunc("/original", apiCfg.GetURLByOriginalURL).Methods("GET")
	router.HandleFunc("/{shortURL}", apiCfg.RedirectURL).Methods("GET")
	router.HandleFunc("/url/{id}", apiCfg.DeleteURL).Methods("DELETE")
	router.HandleFunc("/url/{id}", apiCfg.UpdateURL).Methods("PUT")

}
