package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func FileStreamer(w http.ResponseWriter, r *http.Request) {
	filmDir := "./films"
	filename := r.PathValue("name")
	if filename == "" {
		http.NotFound(w, r)
		return
	}
	fullPath := filepath.Join(filmDir, filename+".mp4")
	fmt.Println(fullPath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeFile(w, r, fullPath)
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	filmDir := "./films"
	validFiles, err := FindValidFiles(filmDir, ".mp4")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	err = json.NewEncoder(w).Encode(validFiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"alive": "true"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
