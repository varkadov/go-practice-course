package storage

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/varkadov/go-practice-course/internal/models"
)

const (
	metricTypeGauge   = "gauge"
	metricTypeCounter = "counter"
)

type MemStorage struct {
	mutex   sync.RWMutex
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		mutex:   sync.RWMutex{},
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (s *MemStorage) Get(metricType, metricName string) (*models.Metrics, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if metricType == metricTypeGauge {
		if v, ok := s.gauge[metricName]; ok {
			return &models.Metrics{
				ID:    metricName,
				MType: metricTypeGauge,
				Value: &v,
			}, nil
		}
	}

	if metricType == metricTypeCounter {
		if v, ok := s.counter[metricName]; ok {
			return &models.Metrics{
				ID:    metricName,
				MType: metricTypeCounter,
				Delta: &v,
			}, nil
		}
	}

	return nil, errors.New("metric doesn't exist")
}

func (s *MemStorage) GetAll() []string {
	l := make([]string, 0)

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for k, v := range s.gauge {
		l = append(l, fmt.Sprintf("%s/%s: %s", metricTypeGauge, k, strconv.FormatFloat(v, 'f', -1, 64)))
	}

	for k, v := range s.counter {
		l = append(l, fmt.Sprintf("%s/%s: %d", metricTypeCounter, k, v))
	}

	return l
}

func (s *MemStorage) Set(metricType, metricName, metricValue string) (*models.Metrics, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if metricType == metricTypeGauge {
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return nil, err
		}

		s.gauge[metricName] = value

		return &models.Metrics{
			ID:    metricName,
			MType: metricTypeGauge,
			Value: &value,
		}, nil
	}

	if metricType == metricTypeCounter {
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return nil, err
		}

		s.counter[metricName] += value
		newValue := s.counter[metricName]

		return &models.Metrics{
			ID:    metricName,
			MType: metricTypeCounter,
			Delta: &newValue,
		}, nil
	}

	return nil, errors.New("metric type doesn't exist")
}
