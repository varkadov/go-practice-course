package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/varkadov/go-practice-course/internal/models"
)

func (h *Handler) MetricsValueHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m := &models.Metrics{}
	err = json.Unmarshal(body, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	metrics, err := h.storage.Get(m.MType, m.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res, err := json.Marshal(metrics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
	w.WriteHeader(http.StatusOK)
}
