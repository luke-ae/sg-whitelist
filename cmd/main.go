package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/dgraph-io/ristretto"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := chi.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "42069"
	}

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 33, // maximum cost of cache (8GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic("error creating cache layer")
	}

	c := NewHTTPClient(os.Getenv("STARGAZE_REST_API_URL"), cache)
	gql := NewGQLClient(os.Getenv("CONSTELLATIONS_GRAPHQL_URL"), c, cache)

	r.Get("/api/v1beta/on_whitelist/{wallet}", gql.WhitelistedCollections)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, Health{Status: "healthy"})
	})

	slog.Info(fmt.Sprintf("starting server on port %v...", port))

	err = http.ListenAndServe(fmt.Sprintf(":%v", port), r)
	if err != nil {
		panic(fmt.Sprintf("error starting server on port %v", port))
	}

}
