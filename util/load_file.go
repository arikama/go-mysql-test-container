package util

import (
	"io/ioutil"
	"strings"
)

func LoadFile(fileFullPath string) (string, error) {
	bytes, err := ioutil.ReadFile(fileFullPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytes)), nil
}
