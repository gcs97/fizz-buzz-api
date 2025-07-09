package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gcs97/fizz-buzz-api/internal/api"
	"github.com/joho/godotenv"
)

const DEFAULT_PORT = ":8080"

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := api.Router()
	logRouter := api.LoggingMiddleware(router)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, logRouter))
}
