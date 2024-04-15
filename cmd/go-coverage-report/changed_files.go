package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func ParseChangedFiles(filename, prefix string, ignored []string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return parseChangedFiles(data, prefix, ignored)
}

func parseChangedFiles(data []byte, prefix string, ignored []string) ([]string, error) {
	var files []string
	err := json.Unmarshal(data, &files)
	if err != nil {
		return nil, err
	}

	coveredFiles := make([]string, 0)
	for _, file := range files {
		pkg := filepath.Dir(file)

		log.Println(file, pkg, ignored)

		if slices.Contains(ignored, pkg) {
			continue
		}
		coveredFiles = append(coveredFiles, file)
	}

	for i, file := range coveredFiles {
		coveredFiles[i] = filepath.Join(prefix, file)
	}

	return coveredFiles, nil
}
