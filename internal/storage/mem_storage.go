package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/varkadov/go-practice-course/internal/models"
)

const (
	metricTypeGauge   = "gauge"
	metricTypeCounter = "counter"
)

type storage interface {
	Store([]byte) error
	Restore() ([]byte, error)
}

type GaugeMetrics map[string]float64

type CounterMetric map[string]int64

type MemStorage struct {
	mutex sync.RWMutex

	gauge   GaugeMetrics
	counter CounterMetric

	storage       storage
	storeInterval int
}

type StorageType struct {
	Gauge   GaugeMetrics
	Counter CounterMetric
}

func NewMemStorage(storage storage, restore bool, storeInterval int) *MemStorage {
	var counterMetrics CounterMetric
	var gaugeMetrics GaugeMetrics

	if restore {
		counterMetrics, gaugeMetrics = restoreFromStorage(storage)
	} else {
		counterMetrics = make(CounterMetric)
		gaugeMetrics = make(GaugeMetrics)
	}

	ms := &MemStorage{
		mutex:         sync.RWMutex{},
		gauge:         gaugeMetrics,
		counter:       counterMetrics,
		storage:       storage,
		storeInterval: storeInterval,
	}

	if storeInterval > 0 {
		go func(interval int) {
			timer := time.NewTicker(time.Duration(interval) * time.Second)
			for range timer.C {
				_ = ms.Flush()
			}
		}(storeInterval)
	}

	return ms
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

		if s.storeInterval == 0 {
			_ = s.Flush()
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

		if s.storeInterval == 0 {
			_ = s.Flush()
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

func (s *MemStorage) Flush() error {
	d := &StorageType{
		Gauge:   s.gauge,
		Counter: s.counter,
	}

	bb, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return s.storage.Store(bb)
}

func restoreFromStorage(storage storage) (CounterMetric, GaugeMetrics) {
	cm := make(CounterMetric)
	gm := make(GaugeMetrics)

	data, err := storage.Restore()
	if err != nil {
		return cm, gm
	}

	st := &StorageType{}

	err = json.Unmarshal(data, st)
	if err != nil {
		return cm, gm
	}

	return st.Counter, st.Gauge
}
