package handlers

type Storage interface {
	GetAll() []string
	Get(metricType, metricName string) (string, error)
	Set(metricType, metricName, metricValue string) error
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}
