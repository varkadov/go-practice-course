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
	storage   Storage
	dbStorage *storage.DBStorage
}

func NewHandler(storage Storage, dbStorage *storage.DBStorage) *Handler {
	return &Handler{
		storage:   storage,
		dbStorage: dbStorage,
	}
}
