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

	addr := flag.String("a", ":8080", "Server address")
	pollInterval := flag.Int64("p", 2, "Pool Interval")
	reportInterval := flag.Int64("r", 10, "Report interval")
	flag.Parse()

	pollTimer := time.NewTicker(time.Duration(*pollInterval) * time.Second)
	reportTimer := time.NewTicker(time.Duration(*reportInterval) * time.Second)

	for {
		select {
		case <-pollTimer.C:
			{
				fmt.Println("It's time for report")

				runtime.ReadMemStats(&m)

				fmt.Printf("Alloc: %v\n", m.Alloc)
			}
		case <-reportTimer.C:
			{
				fmt.Println("It's time for poll")

				res, err := c.Post(*addr+url, "text/plain", nil)
				_ = res.Body.Close()
				if err != nil {
					_ = fmt.Errorf("%v", err)
					return
				}

				fmt.Printf("Status code: %d\n", res.StatusCode)
			}
		}
	}
}
