package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

func PostMetricHandler(store *storage.MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := strings.ToLower(chi.URLParam(r, "metricType"))
		metricName := strings.ToLower(chi.URLParam(r, "metricName"))
		metricValue := chi.URLParam(r, "metricValue")

		if metricType != "gauge" && metricType != "counter" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if metricType == "gauge" {
			value, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			store.Gauge[metricName] = value
		} else if metricType == "counter" {
			value, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			store.Counter[metricName] += value
		}

		w.WriteHeader(http.StatusOK)
	}
}
