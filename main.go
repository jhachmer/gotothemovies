package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /films/{name}", FileStreamer)
	mux.HandleFunc("GET /films", ListMovies)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func ValidateMovieFile(filename string) bool {
	return strings.HasSuffix(filename, ".mp4")
}
