package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

const url = "/update/gauge/10/10"

func main() {
	c := http.Client{}
	m := runtime.MemStats{}

	addr := flag.String("a", "http://localhost:8080", "Server address")
	pollInterval := flag.Int64("p", 2, "Pool Interval")
	reportInterval := flag.Int64("r", 10, "Report interval")
	flag.Parse()

	pollTimer := time.NewTicker(time.Duration(*pollInterval) * time.Second)
	reportTimer := time.NewTicker(time.Duration(*reportInterval) * time.Second)

	defer reportTimer.Stop()
	defer pollTimer.Stop()

	for {
		select {
		case <-pollTimer.C:
			{
				fmt.Println("It's time for poll")

				runtime.ReadMemStats(&m)

				fmt.Printf("Alloc: %v\n", m.Alloc)
			}
		case <-reportTimer.C:
			{
				fmt.Println("It's time for report")

				res, err := c.Post(*addr+url, "text/plain", nil)
				if err != nil {
					_ = fmt.Errorf("%v", err)
					return
				}
				_ = res.Body.Close()

				fmt.Printf("Status code: %d\n", res.StatusCode)
			}
		}
	}
}
