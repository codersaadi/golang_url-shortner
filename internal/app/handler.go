package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saadi925/url-shortner-golang/internal/database"
	shortener "github.com/saadi925/url-shortner-golang/internal/shortner"
)

type ApiConfig struct {
	DB *database.Queries
}
type shortenRequest struct {
	URL string `json:"url"`
}

type updateShortUrl struct {
	URL      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func (apiConfig *ApiConfig) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println("url ", req.URL)
	shortURL, err := shortener.Shorten(apiConfig.DB, req.URL)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	res := shortenResponse{ShortURL: shortURL}
	json.NewEncoder(w).Encode(res)
}

func (apiConfig *ApiConfig) RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	originalURL, err := shortener.GetOriginalURL(apiConfig.DB, shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (apiConfig *ApiConfig) GetURLByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	int32ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	url, err := apiConfig.DB.GetURLByID(r.Context(), int32(int32ID))
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(url)
}

// getURLByOriginalURL
func (apiConfig *ApiConfig) GetURLByOriginalURL(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	originalURL := query.Get("url")
	url, err := apiConfig.DB.GetURLByOriginalURL(r.Context(), originalURL)
	if err != nil {
		fmt.Println("error ", err)
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(url)
}

// getURLs
func (apiConfig *ApiConfig) GetURLs(w http.ResponseWriter, r *http.Request) {
	urls, err := apiConfig.DB.GetURLs(r.Context())
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(urls)
}

// delete a url
func (apiConfig *ApiConfig) DeleteURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	int32ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = apiConfig.DB.DeleteURL(r.Context(), int32(int32ID))
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// update a url by id
func (apiConfig *ApiConfig) UpdateURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	int32ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var req updateShortUrl
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	url, err := apiConfig.DB.UpdateURL(r.Context(), database.UpdateURLParams{
		ID:          int32(int32ID),
		ShortUrl:    req.ShortUrl,
		OriginalUrl: req.URL,
	})
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(url)
}
