package handlers

import (
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"net/http"
)

func RootHandler(s *storage.MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("RootHandler"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
