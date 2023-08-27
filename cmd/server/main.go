package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/internal/config"
	"github.com/varkadov/go-practice-course/internal/handlers"
	"github.com/varkadov/go-practice-course/internal/middlewares"
	"github.com/varkadov/go-practice-course/internal/storage"
)

func main() {
	s := storage.NewMemStorage()
	r := chi.NewRouter()
	c := config.NewConfig()
	h := handlers.NewHandler(s)

	r.Use(middlewares.WithLogging)
	r.Get("/", h.RootHandler)
	r.Post("/value", h.MetricsValueHandler)
	r.Post("/update", h.MetricsUpdateHandler)
	// Deprecated endpoints
	r.Get("/value/{metricType}/{metricName}", h.GetMetricHandler)
	r.Post("/update/{metricType}/{metricName}/{metricValue}", h.PostMetricHandler)

	fmt.Printf("Server running on %s", c.Addr)

	err := http.ListenAndServe(c.Addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
