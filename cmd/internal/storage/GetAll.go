package storage

import "fmt"

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
