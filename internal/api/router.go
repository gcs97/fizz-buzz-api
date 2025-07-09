package api

import (
	"net/http"

	"github.com/gcs97/fizz-buzz-api/internal/api/handler"
)

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/fizz-buzz", handler.FizzBuzzHandler)
	mux.HandleFunc("/api/v1/stats", handler.StatsHandler)
	return mux
}
