package handlers

import (
	"io"
	"net/http"
)

func (h *Handler) RootHandler(w http.ResponseWriter, r *http.Request) {
	l := h.storage.GetAll()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	for _, v := range l {
		_, _ = io.WriteString(w, v+"\n")
	}
}
