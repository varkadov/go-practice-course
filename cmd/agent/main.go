package main

import (
	"fmt"
	"github.com/varkadov/go-practice-course/cmd/agent/config"
	"net/http"
	"runtime"
	"time"
)

const path = "/update/gauge/10/10"

func main() {
	c := http.Client{}
	m := runtime.MemStats{}
	conf := config.NewConfig()

	pollTimer := time.NewTicker(time.Duration(conf.PollInterval) * time.Second)
	reportTimer := time.NewTicker(time.Duration(conf.ReportInterval) * time.Second)

	defer reportTimer.Stop()
	defer pollTimer.Stop()

	url := "http://" + conf.Addr + path

	for {
		select {
		case <-pollTimer.C:
			{
				runtime.ReadMemStats(&m)
			}
		case <-reportTimer.C:
			{
				res, err := c.Post(url, "text/plain", nil)
				if err != nil {
					_ = fmt.Errorf("%v", err)
					return
				}
				_ = res.Body.Close()
			}
		}
	}
}
