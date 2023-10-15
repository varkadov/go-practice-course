package handlers

import (
	"github.com/varkadov/go-practice-course/internal/models"
	"github.com/varkadov/go-practice-course/internal/storage"
)

type Storage interface {
	GetAll() []string
	Get(metricType, metricName string) (*models.Metrics, error)
	Set(metricType, metricName, metricValue string) (*models.Metrics, error)
}

type Handler struct {
	storage    Storage
	sqlStorage *storage.DBStorage
}

func NewHandler(storage Storage, sqlStorage *storage.DBStorage) *Handler {
	return &Handler{
		storage:    storage,
		sqlStorage: sqlStorage,
	}
}
