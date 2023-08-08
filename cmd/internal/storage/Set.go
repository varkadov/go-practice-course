package storage

import (
	"errors"
	"strconv"
)

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
