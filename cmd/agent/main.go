package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/varkadov/go-practice-course/internal/config"
)

func main() {
	c := http.Client{}
	conf := config.NewConfig()

	m := runtime.MemStats{}
	pollCount := 0

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

			paths := []string{
				fmt.Sprintf("/update/gauge/Alloc/%d", m.Alloc),
				fmt.Sprintf("/update/gauge/BuckHashSys/%d", m.BuckHashSys),
				fmt.Sprintf("/update/gauge/Frees/%d", m.Frees),
				fmt.Sprintf("/update/gauge/GCCPUFraction/%f", m.GCCPUFraction),
				fmt.Sprintf("/update/gauge/GCSys/%d", m.GCSys),
				fmt.Sprintf("/update/gauge/HeapAlloc/%d", m.HeapAlloc),
				fmt.Sprintf("/update/gauge/HeapIdle/%d", m.HeapIdle),
				fmt.Sprintf("/update/gauge/HeapInuse/%d", m.HeapInuse),
				fmt.Sprintf("/update/gauge/HeapObjects/%d", m.HeapObjects),
				fmt.Sprintf("/update/gauge/HeapReleased/%d", m.HeapReleased),
				fmt.Sprintf("/update/gauge/HeapReleased/%d", m.HeapReleased),
				fmt.Sprintf("/update/gauge/HeapSys/%d", m.HeapSys),
				fmt.Sprintf("/update/gauge/LastGC/%d", m.LastGC),
				fmt.Sprintf("/update/gauge/Lookups/%d", m.Lookups),
				fmt.Sprintf("/update/gauge/MCacheInuse/%d", m.MCacheInuse),
				fmt.Sprintf("/update/gauge/MCacheSys/%d", m.MCacheSys),
				fmt.Sprintf("/update/gauge/MSpanInuse/%d", m.MSpanInuse),
				fmt.Sprintf("/update/gauge/MSpanSys/%d", m.MSpanSys),
				fmt.Sprintf("/update/gauge/Mallocs/%d", m.Mallocs),
				fmt.Sprintf("/update/gauge/NextGC/%d", m.NextGC),
				fmt.Sprintf("/update/gauge/NumForcedGC/%d", m.NumForcedGC),
				fmt.Sprintf("/update/gauge/NumGC/%d", m.NumGC),
				fmt.Sprintf("/update/gauge/OtherSys/%d", m.OtherSys),
				fmt.Sprintf("/update/gauge/PauseTotalNs/%d", m.PauseTotalNs),
				fmt.Sprintf("/update/gauge/StackInuse/%d", m.StackInuse),
				fmt.Sprintf("/update/gauge/StackSys/%d", m.StackSys),
				fmt.Sprintf("/update/gauge/Sys/%d", m.Sys),
				fmt.Sprintf("/update/gauge/TotalAlloc/%d", m.TotalAlloc),
				// Other metrics
				fmt.Sprintf("/update/counter/PollCount/%d", pollCount),
				fmt.Sprintf("/update/counter/RandomValue/%d", rand.Int63()),
			}

			var wg sync.WaitGroup
			var errCount int

			for _, path := range paths {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					res, err := c.Post(url, "text/plain", nil)

					if err != nil {
						errCount++
						fmt.Printf("Error: %v\n", err)
						return
					}
					_ = res.Body.Close()
				}(host + path)
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
