package util

import (
	"io/fs"
	"path/filepath"
)

type DirFiles struct {
	Name string
}

func FindValidFiles(root, ext string) ([]DirFiles, error) {
	files := make([]DirFiles, 0)
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == ext {
			file := DirFiles{Name: path}
			files = append(files, file)
		}
		return nil
	})
	return files, nil
}

func Map[TValue, TResult any](values []TValue, fn func(TValue) TResult) []TResult {
	result := make([]TResult, len(values))
	for i, value := range values {
		result[i] = fn(value)
	}
	return result
}

func Filter[T any](values []T, fn func(T) bool) []T {
	result := make([]T, 0, len(values))
	for _, value := range values {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}

func Reduce[TValue, TResult any](values []TValue, initialValue TResult, fn func(TResult, TValue) TResult) TResult {
	result := initialValue
	for _, value := range values {
		result = fn(result, value)
	}
	return result
}
