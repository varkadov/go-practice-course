package handlers

import "github.com/varkadov/go-practice-course/internal/models"

type Storage interface {
	GetAll() []string
	Get(metricType, metricName string) (*models.Metrics, error)
	Set(metricType, metricName, metricValue string) (*models.Metrics, error)
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}
