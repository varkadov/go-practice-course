package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/cmd/internal/handlers"
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"github.com/varkadov/go-practice-course/cmd/server/config"
	"log"
	"net/http"
)

func main() {
	s := storage.NewMemStorage()
	r := chi.NewRouter()
	c := config.NewConfig()

	r.Get("/", handlers.RootHandler(s))
	r.Get("/value/{metricType}/{metricName}", handlers.GetMetricHandler(s))
	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.PostMetricHandler(s))

	fmt.Printf("Server running on %s", c.Addr)

	err := http.ListenAndServe(c.Addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
