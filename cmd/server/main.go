package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/cmd/internal/handlers"
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"log"
	"net/http"
)

func main() {
	s := storage.NewMemStorage()
	r := chi.NewRouter()

	r.Get("/", handlers.RootHandler(s))
	r.Get("/value/{metricType}/{metricName}", handlers.GetMetricHandler(s))
	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.PostMetricHandler(s))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
