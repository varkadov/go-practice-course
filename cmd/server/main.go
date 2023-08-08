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

	r.Route("/", func(r chi.Router) {
		r.Route("/update", func(r chi.Router) {
			r.Get("/{metricType}/{metricName}/{metricValue}", handlers.GetMetricHandler(s))
			r.Post("/{metricType}/{metricName}/{metricValue}", handlers.PostMetricHandler(s))
		})
		r.Get("/", handlers.RootHandler(s))
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
