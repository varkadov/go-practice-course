package storage

import (
	"errors"
	"fmt"
)

type MemStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}
}

func (s *MemStorage) Get(t, n string) (string, error) {
	if t == "gauge" {
		if v, ok := s.Gauge[n]; ok {
			return fmt.Sprintf("%f", v), nil
		}
		return "", errors.New("metric doesn't exist")
	}

	if t == "counter" {
		if v, ok := s.Counter[n]; ok {
			return fmt.Sprintf("%d", v), nil
		}
		return "", errors.New("metric doesn't exist")
	}

	return "", errors.New("metric doesn't exist")
}
