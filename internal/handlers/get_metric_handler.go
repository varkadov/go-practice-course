package handlers

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetMetricHandler(w http.ResponseWriter, r *http.Request) {
	metricType := strings.ToLower(chi.URLParam(r, "metricType"))
	metricName := strings.ToLower(chi.URLParam(r, "metricName"))

	m, err := h.storage.Get(metricType, metricName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	v := ""

	if m.Value != nil {
		v = strconv.FormatFloat(*m.Value, 'f', -1, 64)
	} else if m.Delta != nil {
		v = strconv.FormatInt(*m.Delta, 10)
	}

	_, err = io.WriteString(w, v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
