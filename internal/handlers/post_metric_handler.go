package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/internal/storage"
)

func PostMetricHandler(storage *storage.MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := strings.ToLower(chi.URLParam(r, "metricType"))
		metricName := strings.ToLower(chi.URLParam(r, "metricName"))
		metricValue := chi.URLParam(r, "metricValue")

		err := storage.Set(metricType, metricName, metricValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
