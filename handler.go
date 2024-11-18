package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
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

type File struct {
	Name     string
	FileInfo fs.FileInfo
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	filmDir := "./films"
	entries, err := os.ReadDir(filmDir)
	if err != nil {
		http.NotFound(w, r)
	}
	jsonEncoder := json.NewEncoder(w)
	for _, entry := range entries {
		if name := entry.Name(); ValidateMovieFile(name) {
			info, err := entry.Info()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			fi := File{
				Name:     name,
				FileInfo: info,
			}
			err = jsonEncoder.Encode(fi)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
