package main

import (
	"log"
	"movieExample/metadata/internal/controller"
	hand "movieExample/metadata/internal/handler"
	"movieExample/metadata/memory"
	"net/http"
)

const (
	host = "localhost:8080"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.New()
	ctrl := controller.New(repo)
	h := hand.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(host, nil); err != nil {
		panic(err)
	}
}
