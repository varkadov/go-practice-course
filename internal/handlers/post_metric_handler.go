package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) PostMetricHandler(w http.ResponseWriter, r *http.Request) {
	metricType := strings.ToLower(chi.URLParam(r, "metricType"))
	metricName := strings.ToLower(chi.URLParam(r, "metricName"))
	metricValue := chi.URLParam(r, "metricValue")

	m, err := h.storage.Set(metricType, metricName, metricValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := ""

	if m.Value != nil {
		value = strconv.FormatFloat(*m.Value, 'f', -1, 64)
	} else if m.Delta != nil {
		value = strconv.FormatInt(*m.Delta, 10)
	}

	_, err = w.Write([]byte(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
