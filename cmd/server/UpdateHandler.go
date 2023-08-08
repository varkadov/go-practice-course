package main

import (
	"net/http"
	"strconv"
	"strings"
)

func UpdateHandler(store *MemStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		params := strings.Split(r.URL.Path[len("/update/"):], "/")

		if len(params) < 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		metricType := params[0]

		if metricType != "gauge" && metricType != "counter" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		metricName := params[1]
		metricValue := params[2]

		if metricType == "gauge" {
			value, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			store.gauge[metricName] = value
		} else if metricType == "counter" {
			value, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			store.counter[metricName] += value
		}

		w.WriteHeader(http.StatusOK)
	}
}
