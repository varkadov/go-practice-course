package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
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

func main() {
	store := NewMemStorage()
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		params := strings.Split(r.URL.Path[len("/update/"):], "/")

		if len(params) < 2 {
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
			store.counter[metricType] += value
		}

		// Store the metric

		w.WriteHeader(http.StatusOK)

	})

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
