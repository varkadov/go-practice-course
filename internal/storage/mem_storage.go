package storage

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
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

// TODO use sync.RWMutex https://github.com/varkadov/go-practice-course/pull/5#discussion_r1289184589
func NewMemStorage() *MemStorage {
	return &MemStorage{
		mutex:   sync.RWMutex{},
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (s *MemStorage) Get(metricType, metricName string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if metricType == metricTypeGauge {
		if v, ok := s.gauge[metricName]; ok {
			return strconv.FormatFloat(v, 'f', -1, 64), nil
		}
		return "", errors.New("metric doesn't exist")
	}

	if metricType == metricTypeCounter {
		if v, ok := s.counter[metricName]; ok {
			return fmt.Sprintf("%d", v), nil
		}
		return "", errors.New("metric doesn't exist")
	}

	return "", errors.New("metric doesn't exist")
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

func (s *MemStorage) Set(metricType, metricName, metricValue string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if metricType == metricTypeGauge {
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return err
		}
		s.gauge[metricName] = value
		return nil
	}

	if metricType == metricTypeCounter {
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return err
		}
		s.counter[metricName] += value
		return nil
	}

	return errors.New("metric type doesn't exist")
}
