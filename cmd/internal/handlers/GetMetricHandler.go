package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"io"
	"net/http"
	"strings"
)

func GetMetricHandler(storage *storage.MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := strings.ToLower(chi.URLParam(r, "metricType"))
		metricName := strings.ToLower(chi.URLParam(r, "metricName"))

		v, err := storage.Get(metricType, metricName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		_, err = io.WriteString(w, v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
