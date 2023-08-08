package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

const url = "http://localhost:8080/update/gauge/10/10"

func main() {
	c := http.Client{}
	m := runtime.MemStats{}

	pollChan := make(chan interface{})
	reportChan := make(chan interface{})

	go pollFn(pollChan)
	go reportFn(reportChan)

	for {
		select {
		case <-pollChan:
			{
				fmt.Println("It's time for report")

				runtime.ReadMemStats(&m)

				fmt.Printf("Alloc: %v\n", m.Alloc)
			}
		case <-reportChan:
			{
				fmt.Println("It's time for poll")

				res, err := c.Post(url, "text/plain", nil)
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

func pollFn(ch chan interface{}) {
	for {
		time.Sleep(pollInterval)
		ch <- nil
	}
}

func reportFn(ch chan interface{}) {
	for {
		time.Sleep(reportInterval)
		ch <- nil
	}
}
