package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/varkadov/go-practice-course/cmd/internal/handlers"
	"github.com/varkadov/go-practice-course/cmd/internal/storage"
	"log"
	"net/http"
)

func main() {
	s := storage.NewMemStorage()
	r := chi.NewRouter()

	addr := flag.String("a", ":8080", "Server address")
	flag.Parse()

	r.Get("/", handlers.RootHandler(s))
	r.Get("/value/{metricType}/{metricName}", handlers.GetMetricHandler(s))
	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.PostMetricHandler(s))

	fmt.Printf("Server running on %s", *addr)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
