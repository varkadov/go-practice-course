package middlewares

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lw := &loggingResponseWriter{ResponseWriter: w}

		next.ServeHTTP(lw, r)

		log.WithFields(log.Fields{
			"uri":      r.RequestURI,
			"method":   r.Method,
			"duration": time.Since(start),
			"status":   lw.status,
			"size":     lw.size,
		}).Info("Request")
	})
}
