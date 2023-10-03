package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

var mimeTypes = []string{"application/json", "text/html"}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (grw *gzipResponseWriter) Write(b []byte) (int, error) {
	return grw.Writer.Write(b)
}

func WithGzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseContentType := r.Header.Get("Content-Type")
		shouldCompress := false

		for _, mt := range mimeTypes {
			if responseContentType == mt {
				shouldCompress = true
				break
			}
		}

		// Handle requests
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			reader, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error creating gzip reader: %v", err), http.StatusBadRequest)
				return
			}
			defer reader.Close()

			r.Body = reader
			r.Header.Del("Content-Encoding")
		}

		// Handle responses
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") && shouldCompress {
			w.Header().Set("Content-Encoding", "gzip")
			gzw := gzip.NewWriter(w)
			defer gzw.Close()
			w = &gzipResponseWriter{Writer: gzw, ResponseWriter: w}
		}

		next.ServeHTTP(w, r)
	})
}
