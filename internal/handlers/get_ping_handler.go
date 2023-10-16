package handlers

import "net/http"

func (h *Handler) GetPingHandler(w http.ResponseWriter, r *http.Request) {
	err := h.dbStorage.Ping()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
