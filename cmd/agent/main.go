package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/varkadov/go-practice-course/internal/config"
	"github.com/varkadov/go-practice-course/internal/models"
)

func uint64ToFloat64(v uint64) *float64 {
	nv := float64(v)
	return &nv
}

func uint32ToFloat64(v uint32) *float64 {
	nv := float64(v)
	return &nv
}

func uint64ToPointerUint64(v int64) *int64 {
	return &v
}

func main() {
	c := resty.New().SetRetryCount(2).SetRetryWaitTime(2 * time.Second)
	conf := config.NewConfig()

	m := runtime.MemStats{}
	var pollCount int64 = 0

	pollTimer := time.NewTicker(time.Duration(conf.PollInterval) * time.Second)
	reportTimer := time.NewTicker(time.Duration(conf.ReportInterval) * time.Second)

	defer reportTimer.Stop()
	defer pollTimer.Stop()

	rand.New(rand.NewSource(time.Now().UnixNano()))

	host := "http://" + conf.Addr

	for {
		select {
		case <-pollTimer.C:
			pollCount++
		case <-reportTimer.C:
			runtime.ReadMemStats(&m)

			metrics := []models.Metrics{
				{
					ID:    "Alloc",
					MType: "gauge",
					Value: uint64ToFloat64(m.Alloc),
				},
				{
					ID:    "BuckHashSys",
					MType: "gauge",
					Value: uint64ToFloat64(m.BuckHashSys),
				},
				{
					ID:    "Frees",
					MType: "gauge",
					Value: uint64ToFloat64(m.Frees),
				},
				{
					ID:    "GCCPUFraction",
					MType: "gauge",
					Value: &m.GCCPUFraction,
				},
				{
					ID:    "GCSys",
					MType: "gauge",
					Value: uint64ToFloat64(m.GCSys),
				},
				{
					ID:    "HeapAlloc",
					MType: "gauge",
					Value: uint64ToFloat64(m.HeapAlloc),
				},
				{
					ID:    "HeapIdle",
					MType: "gauge",
					Value: uint64ToFloat64(m.HeapIdle),
				},
				{
					ID:    "HeapInuse",
					MType: "gauge",
					Value: uint64ToFloat64(m.HeapInuse),
				},
				{
					ID:    "HeapObjects",
					MType: "gauge",
					Value: uint64ToFloat64(m.HeapObjects),
				},
				{
					ID:    "HeapReleased",
					MType: "gauge",
					Value: uint64ToFloat64(m.HeapReleased),
				},
				{
					ID:    "HeapSys",
					MType: "gauge",
					Value: uint64ToFloat64(m.HeapSys),
				},
				{
					ID:    "LastGC",
					MType: "gauge",
					Value: uint64ToFloat64(m.LastGC),
				},
				{
					ID:    "Lookups",
					MType: "gauge",
					Value: uint64ToFloat64(m.Lookups),
				},
				{
					ID:    "MCacheInuse",
					MType: "gauge",
					Value: uint64ToFloat64(m.MCacheInuse),
				},
				{
					ID:    "MCacheSys",
					MType: "gauge",
					Value: uint64ToFloat64(m.MCacheSys),
				},
				{
					ID:    "MSpanInuse",
					MType: "gauge",
					Value: uint64ToFloat64(m.MSpanInuse),
				},
				{
					ID:    "MSpanSys",
					MType: "gauge",
					Value: uint64ToFloat64(m.MSpanSys),
				},
				{
					ID:    "Mallocs",
					MType: "gauge",
					Value: uint64ToFloat64(m.Mallocs),
				},
				{
					ID:    "NextGC",
					MType: "gauge",
					Value: uint64ToFloat64(m.NextGC),
				},
				{
					ID:    "NumForcedGC",
					MType: "gauge",
					Value: uint32ToFloat64(m.NumForcedGC),
				},
				{
					ID:    "NumGC",
					MType: "gauge",
					Value: uint32ToFloat64(m.NumGC),
				},
				{
					ID:    "OtherSys",
					MType: "gauge",
					Value: uint64ToFloat64(m.OtherSys),
				},
				{
					ID:    "PauseTotalNs",
					MType: "gauge",
					Value: uint64ToFloat64(m.PauseTotalNs),
				},
				{
					ID:    "StackInuse",
					MType: "gauge",
					Value: uint64ToFloat64(m.StackInuse),
				},
				{
					ID:    "StackSys",
					MType: "gauge",
					Value: uint64ToFloat64(m.StackSys),
				},
				{
					ID:    "Sys",
					MType: "gauge",
					Value: uint64ToFloat64(m.Sys),
				},
				{
					ID:    "TotalAlloc",
					MType: "gauge",
					Value: uint64ToFloat64(m.TotalAlloc),
				},
				// Counter metrics
				{
					ID:    "PollCount",
					MType: "counter",
					Delta: &pollCount,
				},
				{
					ID:    "RandomValue",
					MType: "counter",
					Delta: uint64ToPointerUint64(rand.Int63()),
				},
			}

			var wg sync.WaitGroup
			var errCount int

			for _, m := range metrics {
				wg.Add(1)

				m := m
				go func(url string) {
					defer wg.Done()
					body, err := json.Marshal(m)
					if err != nil {
						errCount++
						fmt.Printf("Error: %v\n", err)
						return
					}

					_, err = c.R().
						SetHeader("Content-type", "application/json").
						SetBody(body).
						Post(url)
					if err != nil {
						errCount++
						fmt.Printf("Error: %v\n", err)
						return
					}
				}(host)
			}

			wg.Wait()

			if errCount > 0 {
				fmt.Printf("%d metrics have not been sent\n", errCount)
			} else {
				fmt.Println("All metrics have been sent")
			}
		}
	}
}
