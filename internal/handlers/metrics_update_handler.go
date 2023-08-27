package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/varkadov/go-practice-course/internal/models"
)

func (h *Handler) MetricsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	metrics := models.Metrics{}
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	metricValue := ""

	if metrics.Value != nil {
		metricValue = strconv.FormatFloat(*metrics.Value, 'f', -1, 64)
	} else if metrics.Delta != nil {
		metricValue = strconv.FormatInt(*metrics.Delta, 10)
	}

	updatedMetrics, err := h.storage.Set(metrics.MType, metrics.ID, metricValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(updatedMetrics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	w.WriteHeader(http.StatusOK)
}
