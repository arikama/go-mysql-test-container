package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetFiles(relativePath string) ([]string, error) {
	files, err := ioutil.ReadDir(relativePath)
	if err != nil {
		return nil, err
	}
	wd, _ := os.Getwd()
	results := []string{}
	for _, file := range files {
		fullPath := filepath.Join(wd, relativePath, file.Name())
		results = append(results, fullPath)
	}
	return results, nil
}
