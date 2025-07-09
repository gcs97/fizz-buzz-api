package api

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		slog.Info("request",
			slog.String("method", r.Method),
			slog.String("uri", r.RequestURI),
			slog.String("remote_addr", r.RemoteAddr),
			slog.Duration("duration", time.Since(start)),
		)
	})
}
