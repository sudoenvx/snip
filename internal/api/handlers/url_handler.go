package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sudoenvx/snip/internal/database"
	"github.com/sudoenvx/snip/internal/shortener"
)

type Payload struct {
	Url string `json:"url"`
}

type ShortenedUrl struct {
	OriginalUrl string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	Clicks      int    `json:"clicks"`
}

func CreateShortenUrlHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		payload := new(Payload)

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to read body")
			return
		}

		result, err := shortener.ShortenUrl(payload.Url)
		if err != nil {

		}

		query := "insert into urls (original_url, code) values ($1, $2)"
		command, err := db.Pool.Exec(r.Context(), query, payload.Url, result.Code)
		if err != nil {
			fmt.Fprintf(w, "Failed to create url")
		}

		success := command.Insert()
		if !success {

		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td><td>%d</td></tr>", payload.Url, result.Shorten, 0)

	}
}

func CreateRedirectHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		var originalUrl string

		query := "select original_url from urls where code = $1"
		row := db.Pool.QueryRow(r.Context(), query, code)
		row.Scan(&originalUrl)

		updateClicksQuery := "update urls set clicks = clicks + 1 where code = $1"
		command, err := db.Pool.Exec(r.Context(), updateClicksQuery, code)
		if err != nil {
			fmt.Fprintf(w, "Failed to update clicks")
		}

		command.Insert()

		http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
	}
}

func CreateGetAllUrlsHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// Get all urls
		query := "select original_url, code, clicks from urls"
		rows, err := db.Pool.Query(r.Context(), query)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to get urls from database")
			return
		}

		defer rows.Close()

		urlList := make([]ShortenedUrl, 0)

		for rows.Next() {
			url := new(ShortenedUrl)
			err = rows.Scan(&url.OriginalUrl, &url.ShortCode, &url.Clicks)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to scan row")
				return
			}
			urlList = append(urlList, *url)
		}

		// err = json.NewEncoder(w).Encode(urlList)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	fmt.Fprintf(w, "Failed to encode json")
		// 	return
		// }

		w.WriteHeader(http.StatusOK)
		for _, u := range urlList {
			fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td><td>%d</td></tr>", u.OriginalUrl, u.ShortCode, u.Clicks)
		}
	}
}
