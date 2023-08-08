package main

import (
	"log"
	"net/http"
)

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func main() {
	store := NewMemStorage()
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", UpdateHandler(store))

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
