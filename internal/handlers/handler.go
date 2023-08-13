package handlers

import "github.com/varkadov/go-practice-course/internal/storage"

type Handler struct {
	storage *storage.MemStorage
}

func NewHandler(storage *storage.MemStorage) *Handler {
	return &Handler{
		storage: storage,
	}
}
