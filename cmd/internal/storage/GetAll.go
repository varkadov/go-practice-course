package storage

import (
	"fmt"
	"strconv"
)

func (s *MemStorage) GetAll() []string {
	l := make([]string, 0)

	for k, v := range s.gauge {
		l = append(l, fmt.Sprintf("%s/%s: %s", MetricTypeGauge, k, strconv.FormatFloat(v, 'f', -1, 64)))
	}

	for k, v := range s.counter {
		l = append(l, fmt.Sprintf("%s/%s: %d", MetricTypeCounter, k, v))
	}

	return l
}
