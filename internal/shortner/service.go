package shortener

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/saadi925/url-shortner-golang/internal/database"
)

func Shorten(db *database.Queries, url string) (string, error) {
	// Generate short URL
	shortURL := generateShortURL()

	// Save to repository
	_, err := db.CreateURL(context.Background(), database.CreateURLParams{
		OriginalUrl: url,
		ShortUrl:    shortURL,
	})
	if err != nil {
		fmt.Println("err ", err)
		return "", err
	}
	return shortURL, nil
}
func GetOriginalURL(db *database.Queries, shortURL string) (string, error) {
	url, err := db.GetURLByShortURL(context.Background(), shortURL)
	if err != nil {
		return "", err
	}

	return url.OriginalUrl, nil
}

func generateShortURL() string {
	// Use UUID to generate a unique short URL
	return uuid.New().String()[:8]
}
