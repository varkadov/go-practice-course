package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func WithGzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentEncoding := r.Header.Get("Content-Encoding")

		if strings.Contains(contentEncoding, "gzip") {
			reader, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error creating gzip reader: %v", err), http.StatusBadRequest)
				return
			}
			defer reader.Close()

			r.Body = reader
			r.Header.Del("Content-Encoding")
		}

		next.ServeHTTP(w, r)
	})
}
