package handlers

import (
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"io"
	"net/http"
)

func RootHandler(s *storage.MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l := s.GetAll()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		for _, v := range l {
			io.WriteString(w, v+"\n")
		}
	}
}
