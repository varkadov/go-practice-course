package storage

import (
	"errors"
	"fmt"
	"strconv"
)

func (s *MemStorage) Get(t, n string) (string, error) {
	if t == MetricTypeGauge {
		if v, ok := s.gauge[n]; ok {
			return strconv.FormatFloat(v, 'f', -1, 64), nil
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
