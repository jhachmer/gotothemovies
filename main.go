package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "goto:", log.LstdFlags)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", Chain(HealthHandler, Logging(logger)))
	mux.HandleFunc("GET /films/{name}", Chain(FileStreamer, Logging(logger)))
	mux.HandleFunc("GET /films", Chain(ListMovies, Logging(logger)))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
