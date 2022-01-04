package main

import (
	"cs-fundamentals/service_patterns/domain/adding"
	"cs-fundamentals/service_patterns/domain/storage"
	"net/http"
)

func main() {

	storage := storage.NewStorage()
	adderService := adding.NewService(storage)
	http.HandleFunc("/", adding.Endpoint(adderService))
	http.ListenAndServe(":8080", nil)
}
