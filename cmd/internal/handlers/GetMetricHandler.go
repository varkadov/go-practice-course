package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"io"
	"net/http"
	"strings"
)

func GetMetricHandler(s *storage.MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := strings.ToLower(chi.URLParam(r, "metricType"))
		metricName := strings.ToLower(chi.URLParam(r, "metricName"))

		if metricType == "gauge" {
			if v, ok := s.Gauge[metricName]; ok {
				s := fmt.Sprintf("%f", v)

				_, err := io.WriteString(w, s)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
				return
			}

			http.Error(w, "", http.StatusNotFound)
			return
		} else if metricType == "counter" {
			if v, ok := s.Counter[metricName]; ok {
				s := fmt.Sprintf("%d", v)

				_, err := io.WriteString(w, s)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
				return
			}

			http.Error(w, "", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
	}
}
