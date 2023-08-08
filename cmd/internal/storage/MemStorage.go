package storage

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	MetricTypeGauge   = "gauge"
	MetricTypeCounter = "counter"
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
	if t == MetricTypeGauge {
		if v, ok := s.Gauge[n]; ok {
			return fmt.Sprintf("%f", v), nil
		}
		return "", errors.New("metric doesn't exist")
	}

	if t == MetricTypeCounter {
		if v, ok := s.Counter[n]; ok {
			return fmt.Sprintf("%d", v), nil
		}
		return "", errors.New("metric doesn't exist")
	}

	return "", errors.New("metric doesn't exist")
}

func (s *MemStorage) Set(t, n, v string) error {
	if t == "" || n == "" || v == "" {
		return errors.New("invalid params")
	}

	if t == MetricTypeGauge {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		s.Gauge[n] = value
		return nil
	}

	if t == MetricTypeCounter {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		s.Counter[n] += value
		return nil
	}

	return errors.New("metric type doesn't exist")
}
