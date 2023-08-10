package handlers

import (
	"io"
	"net/http"

	"github.com/varkadov/go-practice-course/internal/storage"
)

func RootHandler(s *storage.MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l := s.GetAll()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		for _, v := range l {
			_, _ = io.WriteString(w, v+"\n")
		}
	}
}
