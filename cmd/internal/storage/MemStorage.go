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
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (s *MemStorage) Get(t, n string) (string, error) {
	if t == MetricTypeGauge {
		if v, ok := s.gauge[n]; ok {
			return fmt.Sprintf("%f", v), nil
		}
		return "", errors.New("metric doesn't exist")
	}

	if t == MetricTypeCounter {
		if v, ok := s.counter[n]; ok {
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
		s.gauge[n] = value
		return nil
	}

	if t == MetricTypeCounter {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		s.counter[n] += value
		return nil
	}

	return errors.New("metric type doesn't exist")
}

func (s *MemStorage) GetAll() []string {
	l := make([]string, 0)

	for k, v := range s.gauge {
		l = append(l, fmt.Sprintf("%s/%s: %f", MetricTypeGauge, k, v))
	}

	for k, v := range s.counter {
		l = append(l, fmt.Sprintf("%s/%s: %d", MetricTypeCounter, k, v))
	}

	return l
}
